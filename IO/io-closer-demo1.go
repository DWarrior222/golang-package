package main

import (
	"bufio"
	"fmt"
	"os"
)

func closerDemo() {
	file, err := os.Open("./wirter.txt")
	if err != nil {
		fmt.Println("err is ", err)
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}

func main() {
	closerDemo()
}
