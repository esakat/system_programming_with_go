package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("waiting 5 second")
	time.Sleep(5 * time.Second)
	fmt.Println("Done!")
}
