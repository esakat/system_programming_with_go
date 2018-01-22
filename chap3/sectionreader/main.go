package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	// 文字列からSectionの部分だけを切り出したReaderを作成し、それをos.Stdoutで書きだしている
	reader := strings.NewReader("Example of io.SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}
