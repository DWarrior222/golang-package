package main

import(
	"fmt"
	"strings"
	"unicode"
)

func field() {
	fmt.Println(strings.Fields(" hello my baby"))
	fmt.Println(strings.FieldsFunc(" hello my baby", unicode.IsSpace))
}

func split() {
	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))
	fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", 2))
}

func splitDemo() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.SplitAfter("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}

func main() {
	// field()
	// split()
	splitDemo()
}