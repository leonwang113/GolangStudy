package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main () {
	//长度
	s := "leon,alvin,john"
	fmt.Println(len(s))

	//strings.Split: 将字符串分割成数组
	var ret = strings.Split(s, ",")
	fmt.Println(ret, reflect.TypeOf(ret))   //[]
	fmt.Println(len(ret))

	//strings.Join： 将述组拼接成字符串
	var new_ret = strings.Join(ret, ",")
	fmt.Println(new_ret, reflect.TypeOf(new_ret))

	//string.Index, strings.LastIndex
	var index = strings.Index(s, "leon")
	fmt.Println(index)   // 未找到返回-1
	var index2 = strings.LastIndex(s, "k")
	fmt.Println(index2)  // 未找到返回-1

	//判断包涵 strings.Contains
	b := strings.Contains(s, "leon")
	fmt.Println(b, reflect.TypeOf(b))

	//判断前后缀 strings.HasPrefix, strings.HasSuffix

	fmt.Println(strings.HasPrefix(s, "leon"))
	fmt.Println(strings.HasSuffix(s, "john"))
}
