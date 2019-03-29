package main

import (
	"fmt"
	"os"
)

func scan() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(a, b, c)
}

func scanf() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Println(a, b, c)
}

func scanln() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	fmt.Println(a, b, c)
}

func sscan() {
	var str string = "aa bb cc"
	var a, b, c string
	fmt.Sscan(str, &a, &b, &c)
	fmt.Println(a, b, c)
}

func sscanf() {
	var str string = "aa,bb,cc"
	var a, b, c string
	fmt.Sscanf(str, "%s,%s,%s", &a, &b, &c)
	fmt.Println(a, b, c)
}

func sscanln() {
	var str string = "aa bb cc"
	var a, b, c string
	fmt.Sscan(str, &a, &b, &c)
	fmt.Println(a, b, c)
}

func fscan() {
	var (
		n, a, b, c int
		err        error
	)

	n, err = fmt.Fscan(os.Stdin, &a, &b, &c)
	if err != nil {
		fmt.Printf("输入正确参数%v个，错误参数原因:%v", n, err)
		return
	}
	fmt.Println(a, b, c)
}

func fscan2() {
	var (
		n, a, b, c int
		err        error
	)

	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("err is ", err)
		panic(err)
	}
	defer file.Close()
	n, err = fmt.Fscan(file, &a, &b, &c)
	if err != nil {
		fmt.Printf("输入正确参数%v个，错误参数原因:%v", n, err)
		return
	}
	fmt.Println(a, b, c)
}

func fscanf() {
	var (
		n       int
		a, b, c string
		err     error
	)

	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("err is ", err)
		panic(err)
	}
	defer file.Close()
	n, err = fmt.Fscanf(file, "%s %s %s", &a, &b, &c)
	if err != nil {
		fmt.Printf("输入正确参数%v个，错误参数原因:%v", n, err)
		return
	}
	fmt.Println(a, b, c)
}

func fscanln() {
	var (
		n, a, b, c int
		err        error
	)

	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("err is ", err)
		panic(err)
	}
	defer file.Close()
	n, err = fmt.Fscan(file, &a, &b, &c)
	if err != nil {
		fmt.Printf("输入正确参数%v个，错误参数原因:%v", n, err)
		return
	}
	fmt.Println(a, b, c)
}

func main() {
	// scan()
	// scanf()
	// scanln()
	// sscan()
	// sscanf()
	// sscanln()
	// fscan()
	// fscan2()
	// fscanf()
	fscanln()
}
