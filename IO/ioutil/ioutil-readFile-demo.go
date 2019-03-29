package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CreateFile() {
	file, err := os.Create("./file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("123456789阿萨德")
}

func readFile() {
	r, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("err is ", err)
		panic(err)
	}
	fmt.Printf("%c \n", r)
}

func writeFile() {
	data := []byte("golang中文网")
	err := ioutil.WriteFile("./file.txt", data, 0666)
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
}

func main() {
	// readFile()
	writeFile()
}
