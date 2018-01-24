package main

import (
	"fmt"
	"time"
)

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}

func main() {
	fmt.Println("start sub()")
	go sub()
	time.Sleep(2 * time.Second)
}

/** 無名関数(クロージャ)を利用版
func main() {
	fmt.Println("start sub()")
	// インラインで無名関数を作ってその場でgoroutineで実行
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}() // すぐ後ろに()がついてるのは即時関数呼び出しのため
	time.Sleep(2 * time.Second)
}
**/
