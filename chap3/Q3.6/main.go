package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	/*
		ここにioパッケージを使ったコードをかく
	*/
	stream = io.MultiReader(
		// SectionReaderで指定した箇所から１バイトを読み込み
		io.NewSectionReader(programming, 5, 1), // A
		io.NewSectionReader(system, 0, 1),      // S
		io.NewSectionReader(computer, 0, 1),    // C
		io.NewSectionReader(programming, 8, 1), // I
		io.NewSectionReader(programming, 8, 1), // I
	)

	io.Copy(os.Stdout, stream)
}

/**
$ go run main.go
ASCII
**/
