// 实现将文件中的数据全部读取（显示在标准输出）
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 使用 io.Reader 接口
// 先获取文件的大小（File 的 Stat 方法），之后定义一个该大小的 []byte，通过 Read 一次性读取
func reader() {
	file, err := os.Open("./wirter.txt")
	if err != nil {
		fmt.Println("err is ", err)
		panic(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	size := fileInfo.Size()

	p := make([]byte, size)
	n, err := file.Read(p)
	if n > 0 {
		fmt.Printf("%s \n", p[:n])
		return
	}
	fmt.Printf("%s \n", p)
}

// readerFrom
func readerFrom() {
	file, err := os.Open("./wirter.txt")
	if err != nil {
		fmt.Println("err is ", err)
		panic(err)
	}
	defer file.Close()

	// bufio.NewWriter 实现了ReaderFrom 接口
	writer := bufio.NewWriter(os.Stdout)
	// func (*Writer) ReadFrom
	// func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
	// ReadFrom实现了io.ReaderFrom。
	writer.ReadFrom(file)
	writer.Flush()
}

func main() {
	reader()
}
