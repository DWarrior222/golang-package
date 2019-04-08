package main

import (
	"fmt"
	"regexp"
)

func compile() {
	reg, err := regexp.Compile("abc")
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
	fmt.Println(reg)
	fmt.Printf("Find result is %s \n", reg.Find([]byte("defabc")))
	fmt.Printf("findall result is %s \n", reg.FindAll([]byte("abcdefabceabc"), 10))
	fmt.Printf("findallindex result is %d \n", reg.FindAllIndex([]byte("abcdefabceabc"), 10))
	fmt.Printf("MatchString result is %v \n", reg.MatchString("abcdefabceabc")) // 是否匹配

	reg = regexp.MustCompile("a.")
	fmt.Printf("FindAllString result is %s \n", reg.FindAllString("abcdefabceabc", 2))
	fmt.Printf("FindAllString result is %s \n", reg.FindAllString("abcdefabceabc", -1))
	fmt.Printf("FindAllStringIndex result is %d \n", reg.FindAllStringIndex("abcdefabceabc", -1))

	reg = regexp.MustCompile("a(b*)c")
	fmt.Printf("FindAllStringSubmatch result is %s \n", reg.FindAllStringSubmatch("abcdefabceabc", -1))
	fmt.Printf("FindAllStringSubmatchIndex result is %d \n", reg.FindAllStringSubmatchIndex("abcdefabceabc", -1))
	fmt.Printf("FindAllSubmatch result is %s \n", reg.FindAllSubmatch([]byte("abcdefabceabc"), -1))           // 返回表达式的所有连续匹配的切片
	fmt.Printf("FindAllSubmatchIndex result is %d \n", reg.FindAllSubmatchIndex([]byte("abcdefabceabc"), -1)) // 返回所有连续匹配位置的切片

	fmt.Printf("FindIndex result is %d \n", reg.FindIndex([]byte("abcdefabceabc"))) // 定义最左侧的匹配位置
	fmt.Printf("FindString result is %s \n", reg.FindString("abcdefabceabc"))       // 定义最左侧的匹配位置
	fmt.Printf("FindStringIndex result is %d \n", reg.FindStringIndex("qabcdefabceabc"))
	reg = regexp.MustCompile("a(x*)b(y|z)c")
	fmt.Printf("FindStringSubmatch result is %q \n", reg.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("FindStringSubmatch result is %q \n", reg.FindStringSubmatch("-abzc-"))
	fmt.Printf("String result is %q \n", reg.String())

	reg = regexp.MustCompile("")
	fmt.Printf("Split result is %q \n", reg.Split("abc", 3))
	fmt.Printf("Split result is %q \n", reg.Split("abcdef", 3))

	reg = regexp.MustCompile("a(.?)")
	fmt.Printf("ReplaceAllString result is %q \n", reg.ReplaceAllString("abcabc", "-"))
	fmt.Printf("ReplaceAll result is %q \n", reg.ReplaceAll([]byte("abcabc"), []byte("-")))

	// FindReaderIndex
	// FindReaderSubmatchIndex
	// MatchReader
}

func compilePOSIX() {
	reg, err := regexp.CompilePOSIX("abc")
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
	fmt.Println(reg)
}

func main() {
	compile()
	// compilePOSIX()
}
