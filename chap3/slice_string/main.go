package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `1行目
2行目
3行目 `

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		// 1行毎に改行で分けて読み込み
		line, err := reader.ReadString('\n')
		// #をつけることで、改行文字などもそのまま出力(結果参考)
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
	/*
		scanner := bufio.NewScanner(strings.NewReader(source))
		for scanner.Scan() {
			fmt.Printf("%#v\n", scanner.Text())
		}

		* 結果
		$ go run main.go
		"1行目"
		"2行目"
		"3行目 "
	*/
}

/**
結果
$ go run main.go
"1行目\n"
"2行目\n"
"3行目 "
**/
