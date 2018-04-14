package main

import (
	"fmt"
	"io"
	"os"

	colorable "github.com/mattn/go-colorable"
	isatty "github.com/mattn/go-isatty"
)

var data = "\033[34m\033[47m\033[4mB\033[31me\n\033[24m\033[30m0S\033[49m\033[m\n"

func main() {
	var stdOut io.Writer
	if isatty.IsTerminal(os.Stdout.Fd()) {
		stdOut = colorable.NewColorableStdout()
	} else {
		stdOut = colorable.NewNonColorable(os.Stdout)
	}
	fmt.Fprintln(stdOut, data)
}
