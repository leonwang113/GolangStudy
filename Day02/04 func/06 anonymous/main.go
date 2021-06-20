package main

import (
	"fmt"
	"reflect"
)

/*func(参数列表)(返回参数列表){
	函数体
}*/


//Go语言不支持在函数内部声明普通函数，只能声明匿名函数
func foo() {
	fmt.Println("foo功能")
	f1 := func(){
		fmt.Println("bar功能")
	}
	f1()
}


func main() {
// 匿名函数可以在使用函数的时候再声明调用
	(func() {
		fmt.Println("leon")
	})()

	var z = func(x,y int) (int) {
		return x + y
	}(1,2)
	fmt.Println(z)

	//也可以将匿名函数作为一个func类型数据赋值给变量
	var f = func(){
		fmt.Println("hello leon")
	}
	fmt.Println(reflect.TypeOf(f))
	f()

	foo()
}
