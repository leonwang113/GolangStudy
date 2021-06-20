package main

import "fmt"

func main() {
	info := map[string]string{"name":"leon","age":"18","gender":"male"}
	delete(info,"gender")
	fmt.Println(info)

}
