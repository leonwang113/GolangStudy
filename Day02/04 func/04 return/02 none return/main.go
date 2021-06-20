package main

import "fmt"

//声明函数时没有定义返回值，函数调用的结果不能作为值使用

func foo(){
	fmt.Printf("Hi, leon!")
	return //不写return默认return为空
}

func main() {
	//ret:=foo() //foo() used as value, but it returns nothing  无返回值不能将调用的结果作为值使用
	foo()
}
