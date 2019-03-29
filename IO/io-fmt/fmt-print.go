package main

import (
	"bufio"
	"fmt"
	"os"
)

func sprint() {
	str := fmt.Sprint("默认格式打印出字符串!")
	fmt.Println(str)
}

func sprintf() {
	str := fmt.Sprintf("Format:%s\n", "格式打印出字符串!")
	fmt.Println(str) // Format:格式打印出字符串!
}

func sprintln() {
	str := fmt.Sprintln("默认格式打印出字符串!")
	fmt.Println(str)
}

func fprint() {
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprint(writer, "测试测试")
	writer.Flush()
}

func fprintf() {
	fmt.Fprintf(os.Stdout, "Format:%s\n", "格式打印!")
}

func fprintln() {
	fmt.Fprintln(os.Stdout, "默认格式加换行打印!")
}

func main() {
	// sprint()
	// sprintf()
	// sprintln()
	// fprint()
	// fprintf()
	fprintln()
}
