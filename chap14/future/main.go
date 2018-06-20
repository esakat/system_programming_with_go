package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(path string) chan string {
	// ファイルを読み込み、その結果を返すFutureを返す
	promise := make(chan string)
	go func() {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("read error %s\n", err.Error())
			close(promise)
		} else {
			promise <- string(content)
		}
	}()
	return promise
}

func printFunc(futureSource chan string) chan []string {
	// 文字列中の関数一覧を返すFutureを返す
	promise := make(chan []string)
	go func() {
		var result []string
		for _, line := range strings.Split(<-futureSource, "\n") {
			if strings.HasPrefix(line, "func ") {
				result = append(result, line)
			}
		}
		promise <- result
	}()
	return promise
}

func main() {
	futureSource := readFile("main.go")
	futureFuncs := printFunc(futureSource)
	fmt.Println(strings.Join(<-futureFuncs, "\n"))
}
