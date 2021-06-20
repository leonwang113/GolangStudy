package main

import "fmt"

func main() {
	//零值，比如int类型的零值是0,string类型的零值是""，引用类型的零值是nil。


/*	//arr:= []int{}
	var arr [] int  //var arr [2] int
	arr[0] = 1
	fmt.Println(arr) //index out of range [0] with length 0*/
	//对于引用类型的变量，我们不光要声明它，还要为它分配内容空间

	//make([]Type,size,cap)

	a := make([]int, 2)
	b := make([]int,2,10)
	fmt.Println(a,b)
	fmt.Println(len(a), len(b))
	fmt.Println(cap(a), cap(b))

	//给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向已经分配好的内存区域，设定开始与结束位置，不会发生内存分配操作
	c := make([]int,5)
	d :=c[0:3]
	fmt.Println(c)
	fmt.Println(d)
	c[0] = 100
	fmt.Println(c)
	fmt.Println(d)
}
