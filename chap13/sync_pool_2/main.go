package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}
	pool.Put("removed 1")
	pool.Put("removed 2")
	// GC呼ぶと消える
	runtime.GC()
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}
