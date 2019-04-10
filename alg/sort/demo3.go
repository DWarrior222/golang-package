package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"n", "b", "q", "z", "a", "w"} // 未排序的切片数据
	sort.Sort(sort.StringSlice(s))
	fmt.Println("Sort: ", s)
	fmt.Println("SearchInts:", sort.SearchStrings(s, "w"))
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("Reverse:", s)
}
