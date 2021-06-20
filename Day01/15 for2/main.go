package main

import "fmt"

func main() {
	// 三要素循环
	//if init;condition;post {}

	for i:=1;i<10; i++{
		fmt.Println(i)
	}
//****************************************

	//for遍历
	 hello := "HelloLeon"
	 for k,v := range hello{
	 	fmt.Printf("%d,%c\n", k,v )
	 }


	 //求 1+ 。。。+100
	 var s = 0
	 for i:=1;i<=100;i++{
	 	s += i
	}
	fmt.Println(s)
}
