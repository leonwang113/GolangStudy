package main

import "fmt"

func main() {
	//开头添加元素
	var a = []int{1,2,3}
	fmt.Printf("slice a: %v\n", a)
	a = append([]int{0}, a...)
	fmt.Println(a)
	a = append([]int{-3,-2,-1}, a...)
	fmt.Println(a)
	fmt.Println()
	//在切片开头添加元素一般都会导致内存的重新分配，而且会导致已有元素全部被复制 1 次


	//任意位置插入元素

	//注意听


	//删除元素
	c := []int{30,31,32,33,34,35,36,37}
	//要删除索引为2的元素
	c = append(c[:2], c[3:]...)
	fmt.Println(c)

}
