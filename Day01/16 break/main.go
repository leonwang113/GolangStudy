package main

import "fmt"

func main() {
	//break 终止循环而执行整个循环语句后面的代码
	for i:=0;i<10;i++{
		if i==8 {
			break
		}
		fmt.Println(":",i)
	}
}
