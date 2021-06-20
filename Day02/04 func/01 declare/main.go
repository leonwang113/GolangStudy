package main

import "fmt"

/*	func 函数名(形式参数) (返回值) {
	函数体
	return 返回值   // 函数终止语句
}*/

func calSum100()  {
	var s = 0
	for i:=1;i<=100;i++{
		s += i
	}
	fmt.Println(s)
}
