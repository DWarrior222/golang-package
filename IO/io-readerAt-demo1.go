package main

import (
	"fmt"
	"io"
	"strings"
)

func ReadAtFrom(reader io.ReaderAt, num int) {
	p := make([]byte, num)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d \n", p, n)
}

func main() {
	reader := strings.NewReader("go语言中文网")
	ReadAtFrom(reader, 6)
}
