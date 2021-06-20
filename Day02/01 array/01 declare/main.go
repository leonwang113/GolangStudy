package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var arrName [num]type
	var names [5]string
	fmt.Println(names, reflect.TypeOf(names))

	var ages [5]int
	fmt.Println(ages, reflect.TypeOf(ages))
}
