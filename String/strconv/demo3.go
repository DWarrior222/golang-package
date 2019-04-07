package main

import(
	"fmt"
	"strconv"
)

func formatint() {
	fmt.Println(strconv.FormatInt(127, 10))
	fmt.Println(strconv.Itoa(127))
}

func main() {
	formatint()
}