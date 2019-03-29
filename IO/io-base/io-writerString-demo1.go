package main

import "os"

func main() {
	file, err := os.Create("./wirter-string.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余")
}
