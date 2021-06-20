package main

import "fmt"

func main() {
	//case 1
	a := []int{12,23,34}
	fmt.Println(len(a))
	fmt.Println(cap(a))
	c:=append(a,45)
	a[0] = 100
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println()

	//case 2
	d := make([]int,3,10)
	fmt.Println(d)
	e := append(d, 45,56)
	fmt.Println(e)
	d[0] = 100
	fmt.Println(d)
	fmt.Printf("%p\n", &d)
	fmt.Println(e)
	fmt.Printf("%p\n", &e)
	fmt.Println()
	//case 3

	l := make([]int,5,10)
	fmt.Println(l)
	v1:= append(l,1)
	fmt.Println(v1)
	//fmt.Printf("%p\n", &l)
	fmt.Printf("%p\n", &v1)
	v2:= append(l,2)
	fmt.Println(v2)
	fmt.Printf("%p\n",&v2)
	fmt.Println(v1)
	fmt.Printf("%p\n", &v1)
	fmt.Println()

	//Aware
	var x = make([]int,5,10)
	var y = append(x, 100)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println()

	//interview
	 arr:= [4]int{10,20,30,40}
	 s1:= arr[0:2]  //[10 20]
	 s2:= s1     //[10 20]
	 s3:= append(append(append(s1,1),2),3)  //[10 20 1 2 3]
	 s1[0] = 1000   //[1000,20]
	 fmt.Println(s2[0])   //使用原来数组？
	 fmt.Println(s3[0])   //超过cap， 生成新数组？

}
