package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	randString := rand.Reader
	randfile, err := os.Create("randfile")
	if err != nil {
		panic(err)
	}
	defer randfile.Close()
	// rand.Readerは無限長のランダム文字列のため
	// CopyNで指定したバイト数のみをファイルに書き込み
	io.CopyN(randfile, randString, 1024)
}

/*

$ ls -l randfile
-rw-r--r--  1 tom_red  staff  1024 Jan 24 08:32 randfile
1024バイトのファイルが生成されている
*/
