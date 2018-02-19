package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is running at localhost:8888")
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1500)
	for {
		// ReadFromで送信元のアドレスを把握して送り返せる
		// TCPと違って「データの終了を探りながら受信」もできないため
		// フレームサイズor期待されるデータ分のバッファを作ってそこにデータをまとめて読み込む方法で対応している
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v: %v\n", remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}
