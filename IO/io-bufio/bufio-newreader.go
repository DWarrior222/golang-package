package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func bytesNewBuffer() {
	buf := []byte("123456")
	fmt.Println(buf)
	b := bytes.NewBuffer(buf)
	var data [6]byte
	b.Read(data[:])
	fmt.Println(data[:])
	fmt.Printf(string(data[:]))
}

func newReader() {
	rb := bytes.NewBuffer([]byte("a string to be read"))
	r := bufio.NewReader(rb)
	var buf [128]byte
	n, _ := r.Read(buf[:])
	fmt.Println(n, string(buf[:]))
}

func newReaderSize() {
	rb := bytes.NewBuffer([]byte("a string to be read"))
	r := bufio.NewReaderSize(rb, 8192)
	var buf [128]byte
	n, _ := r.Read(buf[:])
	fmt.Println(n, string(buf[:n]))
}

func main() {
	// bytesNewBuffer()
	// newReader()
	newReaderSize()
}
