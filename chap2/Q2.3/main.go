package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元データ
	source := map[string]string{
		"Hello": "World",
	}

	/*
			以下追加コード
				- 元データをjsonに変換
				- 変換後のデータを標準出力にだす
		  	- 変換後のデータをgzip圧縮して出力
	*/

	// gzip出力用のファイルを定義
	file, err := os.Create("log.gz")
	if err != nil {
		panic(err)
	}

	// gzip圧縮用の変数を宣言
	gzipper := gzip.NewWriter(file)
	gzipper.Header.Name = "log"

	// 標準出力とgzip圧縮を同時に行う
	writer := io.MultiWriter(gzipper, os.Stdout)

	// 元データをjson変換していく
	// エンコード結果をmultiWriterに投げる
	// これでjson変換結果が標準出力とgzip圧縮にかけられる
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	encoder.Encode(source)

	// 終了時にgzipperを閉じる
	defer gzipper.Close()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

/*
実行結果
$ go run main.go
{
  "Hello": "World"
}
{
  "Hello": "World"
}

$ ls
log.gz  main.go
$ gunzip log.gz
$ ls
log     main.go
$ cat log
{
  "Hello": "World"
}
*/
