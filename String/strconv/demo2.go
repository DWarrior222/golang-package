package main

import(
	"fmt"
	"strconv"
)

func parseint() {
	fmt.Println(strconv.ParseInt("1", 10, 0))
	fmt.Println(strconv.ParseUint("1", 10, 0))
	fmt.Println(strconv.Atoi("1"))
	n, err := strconv.ParseInt("128", 10, 8)
	fmt.Println(n, err)
}

func main() {
	parseint()
}