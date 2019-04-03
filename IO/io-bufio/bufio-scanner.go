package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func caclWord() {
	const input = "this is The Golang Standard Library. \nWelcome you!"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "reading input:", err)
	}
	fmt.Println(count)
}

func scannerText() {
	const input = "http://studygolang.com. \nIt is the home of gophers"
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		// scanner.Scan()	// 下次调用 Scan 会覆盖这次的 token
		fmt.Printf("%s, %s \n", scanner.Text(), scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "reading input:", err)
	}
}

func scannerApply() {
	file, err := os.Create("scanner.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("http://studygolang.com.\nIt is the home of gophers.\nIf you are studying golang, welcome you!")
	// 将文件 offset 设置到文件开头
	file.Seek(0, os.SEEK_SET)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	// caclWord()
	// scannerText()
	scannerApply()
}
