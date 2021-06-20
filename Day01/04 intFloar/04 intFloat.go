package main

import (
	"fmt"
	"reflect"
)

func main() {
	//**********************整型******************
	var x = 10
	fmt.Println(x, reflect.TypeOf(x))


	//**********************进制转化******************
	// 十进制转化
	var a int = 10
	fmt.Printf("%d \n", a)  // 10    占位符%d表示十进制
	fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制
	fmt.Printf("%o \n", a)  // 12    占位符%o表示八进制
	fmt.Printf("%x \n", a)  // a     占位符%x表示十六进制

	// 八进制转化
	var b int = 020
	fmt.Printf("%o \n", b)  // 20
	fmt.Printf("%d \n", b)  // 16
	fmt.Printf("%x \n", b)  // 10
	fmt.Printf("%b \n", b)  // 10000

	// 十六进制转化
	var c = 0x12
	fmt.Printf("%d \n", c)  // 18
	fmt.Printf("%o \n", c)  // 22
	fmt.Printf("%x \n", c)  // 12
	fmt.Printf("%b \n", c)  // 10010

	//**************************浮点型********************
	var f1 = 3.14
	fmt.Println("f1:",f1, reflect.TypeOf(f1))

	var f2 = 3.1415 // float64
	fmt.Printf("f2:%f!\n",f2)
	fmt.Printf("f2:%.5f!\n",f2)
	fmt.Println(reflect.TypeOf(f2))

	// 科学计数 2000 = 2*(10**3) 0.0003 - 3*(10**-5)

	var f3 = 3e2 //  3*(10**2) 科学计数表示的数字皆为浮点类型
	fmt.Println(f3, reflect.TypeOf(f3))

	var f4 =3e-2 // 3* (10**-2)
	fmt.Println(f4,reflect.TypeOf(f4))
}