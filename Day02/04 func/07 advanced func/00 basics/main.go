package main

import "fmt"

/*一个高阶函数应该具备下面至少一个特点：
	将一个或者多个函数作为形参
	返回一个函数作为其结果*/


func addCal(x,y int){
	fmt.Println(x+y)
}

func main() {
	//函数名亦是一个变量
	var a = addCal
	a(2,3)
	fmt.Println(a)
	//fmt.Println(add_cal)

	//结论： 函数参数是一个变量，所以，函数名当然可以作为一个参数传入函数体,也可以作为一个返回值
}
