package main

import (
	"fmt"
	"net"
)

func main() {
	// net.Dialではなく、net.ListenUDP()やnet.DialUDP()が使われることが多い
	// マルチキャストなどのようにUDP特有の機能を使う場合は後者を使う必要がある
	// net.Dialは抽象的な...抽象クラスのようなもの
	conn, err := net.Dial("udp4", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Sending to Server")
	_, err = conn.Write([]byte("Hello from Client"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Receiving from server")
	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received: %s\n", string(buffer[:length]))
}
