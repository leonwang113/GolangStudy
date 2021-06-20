package main

//函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
func cal(x,y int) (sum,sub int){
	sum = x + y
	sub = x - y
	return   //return sum sub
}

func main() {
	cal(2,3)
	//  ????????
}
