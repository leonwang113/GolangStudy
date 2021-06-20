package main

import "fmt"

func main() {
	//可根据新增的 key-value 动态的伸缩，因此不存在固定长度或者最大限制，但是也可以选择标明 map 的初始容量

	//make(map[keyType]valueType, cap)

	map2 := make(map[string]int, 100)
	fmt.Println(map2)
}
