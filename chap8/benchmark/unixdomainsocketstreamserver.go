package benchmark

import (
	"bufio"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

/*
 * domainsocket
 */
func UnixDomainSocketStreamServer() {
	path := filepath.Join(os.TempDir(), "bench-unixdomainsocket-sample")
	os.Remove(path) // 定石
	listener, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}
	for {
		conn, _ := listener.Accept()
		go func() {
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			_, err = httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body: ioutil.NopCloser(
					strings.NewReader("Hello world!\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
