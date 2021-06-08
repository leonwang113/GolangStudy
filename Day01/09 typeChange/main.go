package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//var s = "hello"
	//fmt.Println(s+1)   //不存在暗转换

	var x = 100
	// 将整型转换成字符串类型
	fmt.Println(string(x))  // d

	//stronv
	//将整型转换成字符串类型
	var s1 = strconv.Itoa( 100) //"100"
	fmt.Println(s1, reflect.TypeOf(s1))

	//将字符串转换成整型
	var s2 = "200"
	i,ret := strconv.Atoi(s2)
	fmt.Println(i)
	fmt.Println(ret)
}
