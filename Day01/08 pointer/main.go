package main

import (
	"fmt"
	"reflect"
)

func main () {
	//var x = 10
	//此时 x 是一个整形变量

	//var s = "hello"
	//此时 s 是一个字符串变量

	// 取址符号: &      取值符号: *
	var x int8
	x = 10
	fmt.Println(x)  // 10
	fmt.Println(&x) // 0xc000016072

	//P：定义的一个int8指针类型。 存储的是内存地址
	var p *int8    //int类型的指针类型
	p = &x
	fmt.Println(p)
	fmt.Println(&p)

	// 取值操作 *指针类型变量
	fmt.Println(*p)

    // 存地址的需求时才需要指针
    var s = "hello yuan"
    fmt.Println(&s) //s的地址

    var a *string
    a = &s   //a: 是一个string 类型的指针变量
    fmt.Println(*a)
    fmt.Println(s)
    fmt.Println(reflect.TypeOf(a))

}

