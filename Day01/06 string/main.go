package main

import (
	"fmt"
	"reflect"
)

func main () {
	//创建字符串
	var s1 = "Hello leon!"   //Go只支持双引号表示字符串， 单引号表示字符
	fmt.Println(s1, reflect.TypeOf(s1))

	// 索引操作取字符  数据对象【索引】

	var v = string(s1[8])     //
	fmt.Println(v)

	//切片  数据对象【start: end】 注意： 顾头不顾尾

	fmt.Println(s1[6:10])
	fmt.Println(s1[6:])
	fmt.Println(s1[:4])

	// 字符串拼接

	var s2 = "hello"
	var s3 = "world"
	fmt.Println(2+3)
	fmt.Println(s2+s3)

	// 转义符\       \后跟一个普通符号可以赋予这个普通符号以特殊功能； \后跟一个特殊符号， 转义符会将其变成普通符号
	fmt.Printf( "hi, leon! \n")
	fmt.Printf("hi, world! \n" )

	var path = "c:\\Code\\lesson01\\go.exe"
	fmt.Println(path)

	var s = "let\"s go!"
	fmt.Println(s)

	// (5) 打印多行字符串

	var s4 = `
离离原上草，
一岁一枯荣，
野火烧不尽，
春风吹又生。`
	fmt.Println(s4)

	// (6) 格式化输出
	 var name,age = "alvin", 24
	 fmt.Printf("hello %s, 你的年龄是 %d.", name, age)

	 fmt.Println(reflect.TypeOf(name))
	 fmt.Printf("name 的类型是%T\n", name)

     //%v 占位万能符
     fmt.Printf("name: %v\n", name)
	 fmt.Printf("age: %v\n", age)
}