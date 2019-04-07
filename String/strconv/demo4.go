package main

import(
	"fmt"
	"strconv"
)

func ParseBool() {
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("2"))
	fmt.Println(strconv.FormatBool(true))

	var dst []byte
	fmt.Printf("%s", strconv.AppendBool(dst, true))
}

func parseFloat() {
	fmt.Println(strconv.ParseFloat("111.3451234", 32))
	fmt.Println(strconv.ParseFloat("111.3451234", 32))
	fmt.Println(strconv.FormatFloat(111.3451234, 'g', 3, 32))
	fmt.Println(strconv.FormatFloat(111.3451234, 'e', 3, 32))
	fmt.Println(strconv.FormatFloat(111.3451234, 'b', 3, 32))
}

func main() {
	// ParseBool()
	// parseFloat()
	fmt.Println("This is", strconv.Quote("studygolang.com"), "website")
}