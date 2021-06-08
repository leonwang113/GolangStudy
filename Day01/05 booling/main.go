package main

import (
	"fmt"
	"reflect"
)

func main () {

	var v = 34>2
	fmt.Println(v, reflect.TypeOf(v)) //true bool
}


