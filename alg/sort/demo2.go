package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []float64{5.1, 2.2, 6.0, 3, 1, 4} // 未排序的切片数据
	sort.Sort(sort.Float64Slice(s))
	fmt.Println("Sort: ", s)
	fmt.Println("SearchFloat64s:", sort.SearchFloat64s(s, 2.2))
	sort.Sort(sort.Reverse(sort.Float64Slice(s)))
	fmt.Println("Reverse:", s)
}
