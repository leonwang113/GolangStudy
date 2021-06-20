package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var name []Type

	var names = []string{"zhang3","li4","wang5"}
	fmt.Println(names, reflect.TypeOf(names))
}
