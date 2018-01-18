package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	// MultiWriterは複数のio.Writerを受け取り、それら全てに同時に同じ内容を書き込む
	// 例えば writer := io.MultiWriter(file, os.Stdout, os.Stdout)
	// なんかに変更すればターミナル上で二回出力されるようになるよ
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}

/*
実行するとファイルが生成される
*/
