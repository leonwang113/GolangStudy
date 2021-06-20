package main

import "fmt"

func main() {
	//独立嵌套

	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	//关联嵌套
	for i:=0; i<5; i++{
		for j:=0; j<=i; j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}

}
