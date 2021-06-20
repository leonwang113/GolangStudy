package main

import (
	"fmt"
	"reflect"
)

func main() {

	//简单循环
	count := 0       //初始化语句
	for count < 10 { //条件判断
		fmt.Println(count)
		count++ //步进语句
	}

	//无限循环
	fmt.Printf(`
	1. Normal hit
	2. Super hit
	3. Use tool
	4. Run
`)
	for true {
		var choice int
		fmt.Printf("Please enter your choice:")
		fmt.Scanln(&choice)
		fmt.Println(choice, reflect.TypeOf(choice))

		switch choice {
		case 1:
			fmt.Println("Normal Hit")
		case 2:
			fmt.Println("Super Hit")
		case 3:
			fmt.Println("Use tool")
		case 4:
			fmt.Println("Run")
		default:
			fmt.Println("Wrong input!")
		}
	}

}
