package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 読み込み用のファイルを開く
	oldfile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer oldfile.Close()

	// 書き込み用のファイルを開く
	// 新規作成用にこちらはos.Create()を使う
	newfile, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer newfile.Close()

	buffer := make([]byte, 5)
	_, err := os.Stdin.Read(buffer)
	if err == io.EOF {
		fmt.Println("EOF")
		break
	}

	io.Copy(newfile, string(buffer))
}

/*
第１引数に読み込みファイル、第２ファイルに出力するファイル名を指定してください
$ go run main.go old.txt new.txt
*/
