package main

import "fmt"

//指函数被调用之后，执行函数体中的代码所得到的结果，这个结果通过 return 语句返回。
//return 语句将被调函数中的一个确定的值带回到主调函数中，供主调函数使用

/*func 函数名(形参 形参类型)(返回值类型){
	//  函数体
	return 返回值
}

变量 = 函数(实参)    // return 返回的值赋值给某个变量，程序就可以使用这个返回值了。*/

func addCal(x,y int) int{
	return x+y
}

func main() {
	ret := addCal(1,2)
	fmt.Println(ret)
}
