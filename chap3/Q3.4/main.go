package main

import (
	"archive/zip"
	"bytes"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachement; filename=ascii_sample.zip")

	/**
	そのままQ3.3のコードを流用
	**/
	// zip用のファイルを作成
	zipfile, err := os.Create("ascii_sample.zip")
	if err != nil {
		panic(err)
	}
	defer zipfile.Close()

	// zipWriter構造体を作成
	zipWriter := zip.NewWriter(zipfile)
	defer zipWriter.Close()

	// 存在しないnewfileというファイルデータを与える
	writer, err := zipWriter.Create("newfile")
	// zipに閉じ込められたnewfileに対して、zipWriterのwriterを使って、文字を書き込み
	reader := bytes.NewBufferString("hoge")
	io.Copy(writer, reader)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
