package main

import "fmt"

func main() {
	//现声明再赋值

	//var info map[string]string  //没有默认空间 //panic: assignment to entry in nil map
	info := make(map[string]string)
	info["name"] = "leon"
	info["age"] = "29"
	fmt.Println(info)


	//直接声明赋值

	info2 := map[string]string{"name":"leon","age":"29","gender":"male"}
	fmt.Println(info2)

 }
