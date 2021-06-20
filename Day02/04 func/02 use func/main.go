package main

import "fmt"

func calSum100() {

	var s = 0
	for i:=0;i<=100;i++{
		s += i
	}
	fmt.Println(s)
}

func main() {
	//声明一个函数并不会执行函数内代码，只是完成一个一个包裹的作用
	//真正运行函数内的代码还需要对声明的函数进行调用
	calSum100()
}

