package main

import (
	"bufio"
	"os"
)

func main() {
	/**
	出力結果を貯めておいて、ある程度の分量(デフォだと4096バイト)で書き込むbufio
	**/
	buffer := bufio.NewWriter(os.Stdout)
	// 貯めて
	buffer.WriteString("bufio.Writer ")
	// Flush()を発火させると、書き込み実行
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
