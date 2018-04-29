package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	for _, task := range tasks {
		go func() {
			// goroutineが起動するときループが回りきって
			// 全部のtaskが最後のタスクになってします
			fmt.Println(task)
		}()
	}
	time.Sleep(time.Second)
}
