package main

import "fmt"

func main() {
	//自加操作
	var a = 10
	a++ // a += 1 // a= a+1
	fmt.Println(a)

	//比较运算符  返回布尔值

	fmt.Println(2<=1)
	fmt.Println(0<=1)
	fmt.Println( 0!=1)

	//逻辑运算符 与 或 非
	fmt.Println( 2>1 && 2==3)
	fmt.Println( 2>1 || 2==3)
	fmt.Println( !(2>1 || 2==3))
}
