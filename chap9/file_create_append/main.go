package main

import (
	"io"
	"os"
)

// append版
func append() {
	file, err := os.OpenFile("textfile.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Append ですお\n")
}

func main() {
	append()
}
