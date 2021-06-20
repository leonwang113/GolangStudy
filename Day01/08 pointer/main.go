package main

import (
	"fmt"
	"reflect"
)

func main () {

	// 取址符号: &      取值符号: *
	var x int8
	x = 10
	fmt.Println("x's value",x)  // 10
	fmt.Println("x's address:", &x) // 0xc000016072

	//P：定义的一个int8指针类型。 存储的是内存地址
	var p *int8    //int类型的指针类型
	p = &x
	fmt.Println("p's value",p)
	fmt.Println("p's address",&p)

	// 取值操作 *指针类型变量
	fmt.Println("p address's value", *p)

    // 存地址的需求时才需要指针
    var s = "hello leon"
    fmt.Println("s's address:", &s) //s的地址

    var a *string
    a = &s   //a: 是一个string 类型的指针变量
    fmt.Println(*a)
    fmt.Println(s)
    fmt.Println(reflect.TypeOf(a))

    //指针值引用
	var i = 11
	var j = i

	fmt.Printf("i's address: %p\n",&i)
	fmt.Printf("j's address: %p\n",&j)
	// 当使用等号=将一个变量的值赋给另一个变量时，如 j = i ,实际上是在内存中将 i 的值进行了拷贝，
	//可以通过 &i 获取变量 i 的内存地址。此时如果修改某个变量的值，不会影响另一个。

}

