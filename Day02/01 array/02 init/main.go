package main

import (
	"fmt"
	"reflect"
)

func main() {
	//****************先声明再初始化******************
	//有默认值时就开辟了空间
	var names [5]string  //[     ]
	var ages [5]int  //[0 0 0 0 0]

	names[0] = "zhang3"
	names[1] = "li4"
	names[2] = "wang5"
	names[3] = "zhao6"
	names[4] = "sun7"
	fmt.Println(names)

	ages[0] = 23
	ages[1] = 24
	ages[2] = 25
	ages[3] = 26
	ages[4] = 27
	fmt.Println(ages)
	//****************声明初始化******************
	var names2 =[3]string{"zhang3","li4","wang5"}
	var ages2 = [3]int{23,24,25}
	fmt.Println(names2)
	fmt.Println(ages2)
	//****************[...]不限长度******************
	var names3 = [...]string{"zhang3","li4","wang5"}
	var ages3 = [...]int{23,24,25}
	fmt.Println(names3,reflect.TypeOf(names3))
	fmt.Println(ages3, reflect.TypeOf(ages3))

	//****************索引设置******************
	var names4 =[...]string{0:"张三", 1:"李四"}
	fmt.Println(names4, reflect.TypeOf(names4))
}
