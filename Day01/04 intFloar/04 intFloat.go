package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 10
	fmt.Println(x, reflect.TypeOf(x))


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