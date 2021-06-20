package main

import (
	"fmt"
	"reflect"
)

func main() {
	var names = [...]string{"zhang3","li4","wang5","zhao6","sun7"}
	fmt.Println(names, reflect.TypeOf(names))

	s1:= names[0:3]
	s2:= names[2:5]
	fmt.Println(s1,reflect.TypeOf(s1))
	fmt.Println(s2,reflect.TypeOf(s2))
	s1[2] = "leon"
	fmt.Println(s2)
	fmt.Println(names)

/*	type SliceHeader struct {
		Data uintptr   // 指针，指向底层数组中切片指定的开始位置
		Len int        // 长度，即切片的长度
		Cap int        // 最大长度（容量），也就是切片开始位置到数组的最后位置的长度
	}*/

	arr := [8]int{12,23,34,45,56,67,78,89}
	ss:=arr[2:6]
	fmt.Println(ss)		//[34 45 56 67]
	fmt.Println(len(ss)) //4
	fmt.Println(cap(ss)) // 最大长度6


	var a =[...]int{1,2,3,4,5,6}
	a1:=a[0:3]
	a2:=a[0:5]
	a3:=a[1:5]
	a4:=a[1:]
	a5:=a[:]
	a6:=a[1:2]
	fmt.Printf("a1 len: %d, cap: %d\n", len(a1), cap(a1))
	fmt.Printf("a2 len: %d, cap: %d \n",len(a2), cap(a2))
	fmt.Printf("a3 len: %d, cap: %d \n", len(a3), cap(a3))
	fmt.Printf("a4 len: %d, cap: %d \n", len(a4), cap(a4))
	fmt.Printf("a5 len: %d, cap: %d \n", len(a5), cap(a5))
	fmt.Printf("a6 len: %d, cap: %d \n", len(a6), cap(a6))

//	如果没有数组，则是对最初始的切片构建数组存储！
}
