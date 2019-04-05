package main

import(
	"fmt"
	"strings"
	"bytes"
	"bufio"
)

func demo5() {
	fmt.Println(strings.Index("golang package", "ge"))
	fmt.Println(strings.IndexAny("golang package", "p"))
	fmt.Println(strings.IndexFunc("golang package", func(c rune) bool {
		if c > 'o' {
			return true
		}
		return false
	}))
}

func join() {
	fmt.Println(strings.Join([]string{"name=abc", "age=18"}, "&"))
	fmt.Println(strings.Join([]string{"name=abc"}, "&"))
}
func Repeat() {
	fmt.Println(strings.Repeat("haha", 2))
}
func Replace() {
	fmt.Println(strings.Replace("tell me i really wanna know", "know", "know!", -1))
}

func Replacer() {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	// fmt.Println(r.Replace("This is <b>HTML</b>!"))
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriterSize(wb, 8192)
	r.WriteString(w, r.Replace("This is <b>HTML</b>!"))
	w.Flush()
	fmt.Printf("%s", string(wb.Bytes()))
}

func main() {
	// demo5()
	// join()
	// Repeat()
	// Replace()
	Replacer()
}