package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	io.Copy(os.Stdout, strings.NewReader("Go语言中文网"))

	// io.Copy(os.Stdout, os.Stdin)
	// fmt.Println("Got EOF -- bye")
	io.CopyN(os.Stdout, strings.NewReader("Go语言中文网"), 8)
}
