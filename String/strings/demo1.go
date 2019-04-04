// 是否存在某个字符或者子串
package main

import (
	"fmt"
	"strings"
)

func Contains() {
	fmt.Println(strings.Contains("hello, world", "hello"))
	fmt.Println(strings.Contains("路远同学", "路远"))
}

func ContainsAny() {
	fmt.Println(strings.ContainsAny("hello, world", "ok"))
	fmt.Println(strings.ContainsAny("hello, world", "good"))
	fmt.Println(strings.ContainsAny("hello, world", "gat"))
}

func ContainsRune() {
	fmt.Println(strings.ContainsRune("hello, world", 'e'))
}

func main() {
	// Contains()
	// ContainsAny()
	ContainsRune()
}
