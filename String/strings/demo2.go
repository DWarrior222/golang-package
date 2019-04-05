package main

import(
	"fmt"
	"strings"
)

func count() {
	fmt.Println(strings.Count("success", "c"))
	fmt.Println(strings.Count("hahah", "hah"))
}
func main() {
	count()
}