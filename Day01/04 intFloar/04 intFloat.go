package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 10
	fmt.Println(x, reflect.TypeOf(x))
}