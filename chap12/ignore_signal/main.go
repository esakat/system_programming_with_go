package main

import (
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 最初の5秒間はCtrl + Cを受け付ける
	fmt.Println("Accept Ctrl + C for 5 second")
	time.Sleep(time.Second * 5)

	signal.Ignore(syscall.SIGINT, syscall.SIGHUP)

	// 次の5秒は無視する
	fmt.Println("Ignore Ctrl + C for 5 second")
	time.Sleep(time.Second * 5)
}
