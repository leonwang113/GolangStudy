package main

import "fmt"

func main() {
	//continu 退出当次循环

	for i:=0; i<10; i++{
		if i==8{
			continue
		}
		fmt.Println(":", i)
	}
}
