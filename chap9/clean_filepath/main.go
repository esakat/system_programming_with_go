package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// パスをそのままクリーンにする
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))
	// path/path.go

	// パスを絶対パスに整形
	abspath, _ := filepath.Abs("path/filepath/path_unix.go")
	fmt.Println(abspath)
	// /usr/local/go/src/path/filepath/path_unix.go

	// パスを相対ぱすに整形
	relpath, _ := filepath.Rel("/usr/local/go/src",
		"/usr/local/go/src/path/filepath/path.go")

	fmt.Println(relpath)
	// path/filepath/path.go
}
