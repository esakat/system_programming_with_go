package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}

/**
HTTP/1.1 200 Ok
Date: Sun, 21 Jan 2018 23:24:45 GMT
Server: Apache
X-Frame-Options: SAMEORIGIN
X-UA-Compatible: IE=edge;IE=11;IE=10;IE=9
Expires: 0
Content-Length: 135
Connection: close
Content-Type: text/html

<html>
<head>
<meta http-equiv='refresh' content='1; url=http://ascii.jp/&arubalp=e95690d3-ceae-4a5a-a976-5065e0a815'>
</head>
</html>
**/
