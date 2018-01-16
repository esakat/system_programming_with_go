package main

import "os"

func main() {
	os.Stdout.Write([]byte("os.Stdout exmaple\n"))
}
