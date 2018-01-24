package main

import (
	"crypto/rand"
	"os"
)

/*
Q3.2を流用
*/
func main() {
	randString := rand.Reader
	randfile, err := os.Create("randfile")
	if err != nil {
		panic(err)
	}
	defer randfile.Close()

	// これ用に作成した構造体を生成
	copyer := new(Copyer)
	copyer.CopyN(randfile, randString, 1024)
}
