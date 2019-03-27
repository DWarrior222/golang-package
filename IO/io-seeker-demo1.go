package main

import (
	"fmt"
	"io"
	"strings"
)

func seeker() {
	reader := strings.NewReader("GO语言中文网")
	reader.Seek(-6, io.SeekEnd)
	r, size, err := reader.ReadRune()
	fmt.Printf("%c %v %v\n", r, size, err)
}

func main() {
	seeker()
}
