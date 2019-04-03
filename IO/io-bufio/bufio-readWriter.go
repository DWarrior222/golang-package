package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func newReadWriter() {
	rb := bytes.NewBuffer([]byte("a string to be read"))
	wb := bytes.NewBuffer(nil)
	r := bufio.NewReader(rb)
	w := bufio.NewWriter(wb)
	rw := bufio.NewReadWriter(r, w)
	var rbuf [128]byte
	if n, err := rw.Read(rbuf[:]); err != nil {
		panic(err)
		return
	} else {
		fmt.Println(string(rbuf[:n]))
	}
	if _, err := rw.Write([]byte("a string to be written")); err != nil {
		return
	}
	rw.Flush()
	fmt.Println(string(wb.Bytes()))

}

func main() {
	newReadWriter()
}
