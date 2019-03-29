package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	b, err := ioutil.TempDir("/Users/guanshuaijie/Desktop/Personalfolder/", "go-build")
	if err != nil {
		fmt.Println("err is ", err)
		panic(err)
	}
	fmt.Println("b", b)
	f, err := ioutil.TempFile("/Users/guanshuaijie/Desktop/Personalfolder/", "gofmt")
	if err != nil {
		panic(err)
	}
	fmt.Println("f", f)

	defer func() {
		f.Close()
		os.Remove(f.Name())
		os.Remove(b)
	}()
}
