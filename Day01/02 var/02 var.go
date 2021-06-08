package main
import "fmt"

func main() {
	// ***************************************变量的声明***************************************
	var x int8
	var y string
	var z bool
	// 基本数据类型当变量只声明未赋值， 默认为零值
	fmt.Println(x) //0
	fmt.Println(y) //""
	fmt.Println(z) // false

	// 一行声明多个变量
	var a,b int
	fmt.Println(a,b)
	// 多行声明多个变量

	var (
		c int
		d string
		e bool
	)
	fmt.Println(c,d,e)

	// ***************************************变量赋值***************************************

	// 1 现声明再赋值
	var x1 int8
	x1 = 100
	fmt.Println(x1)

	// 2. 声明并赋值
	// var y1 06 string = "yuan"
	var y1 = "yuan"
	fmt.Println(y1)

	// 3. 简洁赋值
	z1 := 60
	fmt.Println(z1)

	// 4。 多重赋值
	var a1,b1 int
	a1,b1 = 2,3
	fmt.Println(a1,b1)

	// 5. 一行赋值多个变量
	var c1,d1,e1 = 100,"hi yuan", false
	fmt.Println(c1,d1,e1)

	// 6. 多行赋值多个变量
	var (
		c2=99
		d2="leon"
		e2= true
	)
	fmt.Println(c2,d2,e2)

	//***************************************匿名变量**************************************
	var _,_,z3 = 1, 2, 3
	fmt.Println(z3)
	//***************************************变量命名**************************************

}
