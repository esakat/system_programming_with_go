package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("waiting 5 seconds")
	after := time.After(5 * time.Second)
	<-after
	fmt.Println("Done!")
}
