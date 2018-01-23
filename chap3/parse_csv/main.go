package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var csvSource = `13101,"100  ","1000003","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾄﾂﾊﾞｼ(１ﾁｮｳﾒ)","東京都","千代田区","一ツ橋(1丁目)",1,0,1,0,0,0
13101,"101  ","1000003","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾄﾂﾊﾞｼ(２ﾁｮｳﾒ)","東京都","千代田区","一ツ橋(2丁目)",1,0,1,0,0,0
`

func main() {
	reader := strings.NewReader(csvSource)
	csvReader := csv.NewReader(reader)
	for {
		// csvReaderのReadメソッドを呼ぶと、行の`,`区切りの文字列配列が返される
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[2], line[6:9])
	}
}
