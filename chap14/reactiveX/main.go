package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {
	// observalを作成
	emitter := make(chan interface{})
	source := observable.Observable(emitter)

	// イベントを受け取るobseverを作成
	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			line := item.(string)
			if strings.HasPrefix(line, "func ") {
				fmt.Println(line)
			}
		},
		ErrHandler: func(err error) {
			fmt.Printf("Encountered error: %v\n", err)
		},
		DoneHandler: func() {
			fmt.Println("Done!")
		},
	}

	// observableとobserverを継続
	sub := source.Subscribe(watcher)

	go func() {
		content, err := ioutil.ReadFile("main.go")
		if err != nil {
			emitter <- err
		} else {
			for _, line := range strings.Split(string(content), "\n") {
				emitter <- line
			}
		}
		close(emitter)
	}()
	<-sub
}
