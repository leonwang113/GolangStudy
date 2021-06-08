package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main () {
	s := "yuan,alvin,john"
	fmt.Println(len(s))

	//strings.Split: 将字符串分割成数组
	var ret = strings.Split(s, ",")
	fmt.Println(ret, reflect.TypeOf(ret))   //[]
	fmt.Println(len(ret))

	//strings.Join： 将述组拼接成字符串
	var new_str = strings.Join()

}
