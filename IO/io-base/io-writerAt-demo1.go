package main

import (
	"fmt"
	"io"
	"os"
)

func writerAtFrom(writer io.WriterAt) {
	n, err := writer.WriteAt([]byte("Go语言中文网"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func main() {
	file, err := os.Create("./wirter.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余")
	writerAtFrom(file)
}
