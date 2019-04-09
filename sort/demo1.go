package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // 未排序的切片数据
	sort.Ints(s)
	fmt.Println("Ints:", s)
	fmt.Println("SearchInts:", sort.SearchInts(s, 2)) // SearchInts()的使用条件为：切片a已经升序排序
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println("Reverse:", s)
}
