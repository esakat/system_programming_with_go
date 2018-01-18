package main

import (
	"net/http"
	"os"
)

/**
	net/httpのRequest構造体は、用途の限定されたio.Writer実装メソッドを持つ
	HTTPリクエストを扱う際の構造体だよ
**/
func main() {
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write(os.Stdout)
}
