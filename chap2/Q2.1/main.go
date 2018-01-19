package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "数値: %d 文字列: %s 浮動小数点: %f", 1, "hoge", 0.131213)
}

/*
結果
$ go run main.go
数値: 1 文字列: hoge 浮動小数点: 0.131213
*/
