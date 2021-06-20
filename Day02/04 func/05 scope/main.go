package main

import "fmt"

//局部变量 ：写在{}中或者函数中或者函数的形参, 都是局部变量
//全局变量 ：函数外面的就是全局变量

func foo () {
	y = 10
	fmt.Println(y)
}

var y = 30


func main() {
	//var y = 20   //fmt.Println(y) //20
	foo()
	fmt.Println(y)
	fmt.Println()


	//if，for语句具备独立开辟作用域的能力
	if true{
		x:=10
		fmt.Println(x)
	}
	//fmt.Println(x)  // undefined: x

	// for的局部空间
	for i:=0;i<10 ;i++  {
		fmt.Println(i)
	}
	//fmt.Println(i)  //undefined: i



}
