package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		// リクエストを受け付けたら、実行される
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// 非同期で処理を行う
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			// リクエスト読み込み
			request, err := http.ReadRequest(
				bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			// リクエストの取り出し
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}
			// デバッグ出力
			fmt.Println(string(dump))
			// レスポンスを作成
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body: ioutil.NopCloser(
					strings.NewReader("Hello, World!\n"),
				),
			}
			// レスポンス書き込み
			response.Write(conn)
			conn.Close()
		}()
	}
}
