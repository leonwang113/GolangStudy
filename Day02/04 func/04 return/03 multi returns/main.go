package main

import "fmt"

func getNameAge()(string, int){
	return "leon",18
}


func main() {
	a,b := getNameAge()
	fmt.Println(a,b)
}
