package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// フォーマットを整えて出力
	// %v はなんでも表示できるフォーマット指定子、 どんな型でもString()メソッドを持っていれば、それを使って表示してくれる
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}
