package main

import (
	"fmt"
)

type WebSite struct {
	Name string
}

func main() {
	site := WebSite{
		Name: "studygolang",
	}
	fmt.Printf("%v \n", site)
	fmt.Printf("%+v \n", site)
	fmt.Printf("%#v \n", site)
	fmt.Printf("%T \n", site)
	fmt.Printf("%t \n", true)
	fmt.Printf("%b \n", 5)
	fmt.Printf("%c \n", 0x4e2d)
	fmt.Printf("%d \n", 0x12)
	fmt.Printf("%o \n", 10)
	fmt.Printf("%q \n", 0x4e2d)
	fmt.Printf("%x \n", 13)
	fmt.Printf("%X \n", 13)
	fmt.Printf("%U \n", 0x4e2d)
	fmt.Printf("%s \n", []byte("Go语言中文网"))
	fmt.Printf("%q \n", "GO语言中文网")
	fmt.Printf("%x \n", "golang")
	fmt.Printf("%X \n", "golang")
	fmt.Printf("%p \n", &site)
}
