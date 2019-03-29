package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	reader := strings.NewReader("asdfghjkl;")
	res, err := ioutil.ReadAll(reader)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("%c", res)
}
