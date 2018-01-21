package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP1.0\r\nHost: ascii.jp\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// ヘッダーを表示
	fmt.Println(res.Header)
	// ボディを表示
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

/*
実行結果
map[Content-Type:[text/html] X-Frame-Options:[SAMEORIGIN] X-Ua-Compatible:[IE=edge;IE=11;IE=10;IE=9] Content-Length:[135] Date:[Sun, 21 Jan 2018 23:35:54 GMT] Server:[Apache] Expires:[0]]
<html>
<head>
<meta http-equiv='refresh' content='1; url=http://ascii.jp/&arubalp=e95690d3-ceae-4a5a-a976-5065e0a815'>
</head>
</html>
*/
