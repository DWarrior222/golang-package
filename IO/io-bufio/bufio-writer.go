package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func newWriterSize() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriterSize(wb, 8192)
	w.Write([]byte("hello"))
	w.Write([]byte(", world"))
	fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
	w.Flush()
	fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
}

func buffered() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriterSize(wb, 8192)
	fmt.Println(w.Buffered())
	// w.WriteByte('a')
	w.Write([]byte("sdfghj"))
	fmt.Println(w.Buffered())
	w.Flush()
	fmt.Println(w.Buffered())
}

func available() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriterSize(wb, 8192)
	fmt.Println(w.Available())
	w.Write([]byte("sdfghj"))
	fmt.Println(w.Available())
}
func main() {
	// newWriterSize()
	// buffered()
	available()
}
