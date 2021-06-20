package main

import (
	"fmt"
	"strings"
)

func main() {
	//数据类型从存储方式分为两类：值类型和引用类型

	//***************值类型*********************
	//(int,float,bool,string),array, struct 都属于值类型
	//特点：变量直接存储值， 内存通常在栈中分配， 栈在函数调用完会被释放。
	//不管是否已经赋值， 编译器为其分配内存， 此时该值存储于栈上

	var a int
	var b string
	var c bool
	var d [2]int
	fmt.Println(a, &a) //0 0xc000016078
	fmt.Println(b, &b) // 0xc000010240
	fmt.Println(c, &c) //false 0xc000016080
	fmt.Println(d, &d) //[0 0] &[0 0]

	//整形赋值
	var e = 10
	f := e
	f = 101
	fmt.Printf("e: %v, e的内存地址是%p\n", e, &e)
	fmt.Printf("f:%v, f的内存地址是%p\n", f, &f)
	//数组赋值
	var g = [3]int{1, 2, 3}
	h := g
	h[1] = 1000
	fmt.Printf("g: %v, g的内存地址是%p\n", g, &g)
	fmt.Printf("h: %v, h的内存地址是%p\n", h, &h)

	//字符串处理
	s := "hello world"
	ns := strings.ToUpper(s)

	fmt.Printf("s: %v, s的内存地址是%p\n", s, &s)
	fmt.Printf("ns: %v, ns的内存地址是%p \n", ns, &ns)

	//***************引用类型*********************

	//指针，slice, map, chan, interface 等都是应用类型
	//特点：变量存储的是一个地址，这个地址存储最终的值

}
