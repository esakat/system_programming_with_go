package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	// Lock() / Unlock() をペアで呼び出してロックする
	mutex.Lock()
	id++
	mutex.Unlock()
	return id
}
func main() {
	var mutex sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
		}()
	}
}
