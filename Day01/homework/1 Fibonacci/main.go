package main

import "fmt"

func main() {
/*	斐波那契数列指的是一个数列 0, 1, 1, 2, 3, 5, 8, 13,
	                      0  1  2  3  4  5  6  7
	特别指出：第0项是0，第1项是第一个1。从第三项开始，每一项都等于前两项之和。*/
	var f = 0
	for i:=0;i<=7;i++ {
		f += i
		fmt.Println(i,":",f)
	}

}
