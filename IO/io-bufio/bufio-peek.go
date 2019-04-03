package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"time"
)

func peek() {
	rb := bytes.NewBuffer([]byte("123456789"))
	r := bufio.NewReader(rb)
	b1, _ := r.Peek(5)
	fmt.Println(string(b1))
	b2, err := r.Peek(11)
	if err != nil {
		fmt.Println("err is ", err)
	}
	fmt.Println(string(b2))
}

func peek2Main() {
	// rb := bytes.NewBuffer([]byte("http://studygolang.com.\t It is the home of gophers"))
	rb := strings.NewReader("http://studygolang.com.\t It is the home of gophers")
	r := bufio.NewReaderSize(rb, 14)
	go peek2(r)
	go r.ReadBytes('\t')
	time.Sleep(1e8)
}

func peek2(reader *bufio.Reader) {
	line, _ := reader.Peek(14)
	fmt.Printf("%s\n", line)
	// time.Sleep(1)
	fmt.Printf("%s\n", line)
}

func main() {
	// peek()
	peek2Main()
}
