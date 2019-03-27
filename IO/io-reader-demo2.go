package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReaderExample() {
FOREND:
	for {
		readerMenu()
		var ch string
		fmt.Scanln(&ch)
		var (
			data []byte
			err  error
		)
		switch strings.ToLower(ch) {
		case "1":
			fmt.Println("请输入不多于9个字符，按回车结束")
			data, err = readFrom(os.Stdin, 11)
		case "2":
			file, err := os.Open("./test.txt")
			if err != nil {
				fmt.Println("err is ", err)
				continue
			}
			data, err = readFrom(file, 9)
			file.Close()
		case "3":
			data, err = readFrom(strings.NewReader("tell me what you are think about i really wanna know"), 12)
		case "4":
			fmt.Println("未实现")
		case "b":
			fmt.Println("返回上级菜单")
			break FOREND
		case "q":
			fmt.Println("程序退出")
			os.Exit(0)
		default:
			fmt.Println("输入错误")
			continue
		}

		if err != nil {
			fmt.Println("数据读取失败，可以试试从其他输入源读取！")
		} else {
			fmt.Printf("读取到的数据是：%s\n", data)
		}
	}
}

func readFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	// reader.Read会将 len(p) 个字节写入p中，并返回读取的字节数n， 0 <= n <= len(p)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func readerMenu() {
	fmt.Println("")
	fmt.Println("*******从不同来源读取数据*********")
	fmt.Println("*******请选择数据源，请输入：*********")
	fmt.Println("1 表示 标准输入")
	fmt.Println("2 表示 普通文件")
	fmt.Println("3 表示 从字符串")
	fmt.Println("4 表示 从网络")
	fmt.Println("b 返回上级菜单")
	fmt.Println("q 退出")
	fmt.Println("***********************************")
}

func main() {
	ReaderExample()
}
