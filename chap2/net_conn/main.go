package main

import (
	"io"
	"net"
	"os"
)

func main() {
	// net.Connはio.Writerのインタフェース
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r^nHost: ascii.jp\r\n\r\n")

	// レスポンスを画面に出力する処理 3章で詳細説明
	io.Copy(os.Stdout, conn)

	/*
		// HTTPリクエストの作成も可能
		req, err := http.NewRequest("GET", "http://ascii.jp", nil)
		// Writeメソッドを使うことで、io.Writerに書き出せる
		// connを渡せばそのままサーバにリクエストを飛ばせる
		req.Write(conn)
	*/
}
