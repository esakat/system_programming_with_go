package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// data変数に格納されたビッグエンディアンの10000という値が、他のエンディアン環境で10000で表示される
func main() {
	// 32ビットのビッグエンディアンデータ(10000)
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	// エンディアン変換
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}
