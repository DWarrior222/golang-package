package main

import(
	"fmt"
	"strings"
)

func demo() {
	fmt.Println(strings.HasPrefix("golang package", "go"))
	fmt.Println(strings.HasSuffix("golang package", "ge"))
}

func main() {
	demo()
}