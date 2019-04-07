package main

import(
	"fmt"
	"bytes"
)

func contains() {
	fmt.Println(bytes.Contains([]byte("1324567"), []byte("1")))
}

func main() {
	contains()
}