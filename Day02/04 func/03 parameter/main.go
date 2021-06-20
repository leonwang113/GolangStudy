package main

import "fmt"

//为了能够在函数调用的时候动态传入一些值给函数，就有了参数的概念

/*	// 函数声明
	func 函数名(形式参数1 参数1类型,形式参数2 参数2类型,...){
	函数体
	}
	// 调用函数
	函数名(实际参数1,实际参数2,...)*/

//函数每次调用可以传入不同的实际参数，传参的过程其实就是变量赋值的过程，将实际参数按位置分别赋值给形参

func calSum(n int) {

	var s = 0
	for i := 1; i <= n; i++ {
		s += i
	}
	fmt.Println(s)
}

//位置参数, 也称必备参数，指的是必须按照正确的顺序将实际参数传到函数中
func addCal(x int, y int){
	fmt.Println(x+y)
}

//不定长参数 ...

func sum(nums ...int)  {
	fmt.Println("nums",nums)
	total := 0
	for _,nums := range nums {
		total += nums
		//fmt.Println(i)  //index
	}
	fmt.Println(total)
}

//可变参数通常要作为函数的最后一个参数
func sumBase(base int, nums ...int) int {
	fmt.Println(base, nums)
	sum := base
	for _, v := range nums {
		sum = sum + v
	}
	return sum
}


func main() {

	calSum(100)
	calSum(200)
	calSum(1000)
	addCal(100,200)
	sum(11,22,33)
	sum(10,20,30,40,50,60,70)

	ret := sumBase(10,2,3,4)
	fmt.Println(ret)

}