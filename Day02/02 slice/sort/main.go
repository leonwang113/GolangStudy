package main

import (
	"fmt"
	"sort"
)

func main() {
	a:=[]int{10,2,3,100}
	sort.Ints(a)
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a[:])))
	fmt.Println(a)
}
