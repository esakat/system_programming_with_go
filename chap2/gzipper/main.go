package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	/*
		受け取ったデータをgzip圧縮して
		上で宣言されたfileに中継する
		これ以外にもハッシュ値の計算とかもあるらしい
	*/
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()
}
