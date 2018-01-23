package main

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
)

func main() {
	// zip用のファイルを作成
	zipfile, err := os.Create("test.zip")
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
