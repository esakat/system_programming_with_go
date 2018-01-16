package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}

/*
Writeされた内容をバッファに記録しておいて
次の行で呼び出しているよ
*/
