package main

import "fmt"

func main() {
	//通过key访问值

	info:= map[string]string{"name":"leon","age":"29","gender":"male"}
	val,isExist := info["name"]

	if isExist{
		fmt.Println(val)
		fmt.Println(isExist)
		fmt.Println()
	}else{
		fmt.Println("key not exist")
	}


	//循环访问所有键值
	for k,v := range info{
		fmt.Println(k,v)
	}
}
