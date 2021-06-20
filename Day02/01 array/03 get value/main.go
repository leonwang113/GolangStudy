package main

import "fmt"

func main() {

	var names = [...] string{"zhwang3","li4","wang5","zhao6","sun7"}
	//索引取值
	fmt.Println(names[2])

	//切片取值
	fmt.Println(names[0:4])
	fmt.Println(names[0:])
	fmt.Println(names[:3])
	//循环取值1
	for i:=0;i<len(names);i++{
		fmt.Println(i,names[i])
	}
	//循环取值2
	for k,v := range names{
		fmt.Println(k,v)
	}
}
