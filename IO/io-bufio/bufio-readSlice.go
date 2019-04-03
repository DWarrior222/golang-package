package main

import (
	"bufio"
	"fmt"
	"strings"
)

func readSlice() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadSlice('\n')
	fmt.Printf("this line: %s", line)
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("this line: %s\n", line)
	fmt.Printf("this n: %s\n", n)
}

func readBytes() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadBytes('\n')
	fmt.Printf("this line: %s", line)
	n, _ := reader.ReadBytes('\n')
	fmt.Printf("this line: %s\n", line)
	fmt.Printf("this n: %s\n", n)
}

func readString() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadString('\n')
	fmt.Printf("this line: %s", line)
	n, _ := reader.ReadString('\n')
	fmt.Printf("this line: %s\n", line)
	fmt.Printf("this n: %s\n", n)
}

func main() {
	// readSlice()
	// readBytes()
	readString()
}
