package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func demo1() {
	var ch byte
	fmt.Scanf("%c\n", &ch)

	buffer := new(bytes.Buffer)
	err := buffer.WriteByte(ch)
	if err == nil {
		fmt.Println("写入一个字节成功")
		newCh, _ := buffer.ReadByte()
		fmt.Printf("读到的字节 %c\n", newCh)
	} else {
		fmt.Println("写入错误", err)
	}
}

func demo2() {
	var ch byte
	fmt.Scanf("%c\n", &ch)

	writer := bufio.NewWriter(os.Stdout)
	err := writer.WriteByte(ch)

	if err == nil {
		fmt.Println("写入一个字节成功")
	} else {
		fmt.Println("写入错误", err)
	}
}

func demo3() {
	reader := bufio.NewReader(os.Stdin)
	newCh, _ := reader.ReadByte()
	fmt.Printf("读到的字节 %c\n", newCh)
}

func main() {
	// demo1()
	// demo2()
	demo3()
}
