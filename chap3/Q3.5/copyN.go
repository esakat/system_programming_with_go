package main

import (
	"io"
)

type Copyer struct {
}

func (c Copyer) CopyN(dest io.Writer, src io.Reader, length int64) {
	// LimitReaderで先頭から指定したバイト数だけ読み込む
	reader := io.LimitReader(src, length)
	io.Copy(dest, reader)
}
