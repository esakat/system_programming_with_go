package main

import (
	"fmt"
	"time"
)

func main() {
	// 2秒を測る
	var t (time.Duration) = 2

	fmt.Println("start time is : ", time.Now())
	timer := time.After(t * time.Second)
	dispalyTime := <-timer
	fmt.Println("end time is   : ", dispalyTime)
}

/*
$ go run main.go
start time is :  2018-01-29 09:04:16.944122 +0900 JST
end time is   :  2018-01-29 09:04:18.946796 +0900 JST
*/
