package main

import "fmt"

func main() {
	var user string
	var pwd int

	fmt.Println("Please input username:")
	fmt.Scanln(&user)
	fmt.Println("Please input password:")
	fmt.Scanln(&pwd)
	if user == "leon" && pwd == 123 {
		var score int
		fmt.Println("Please input score")
		fmt.Scanln(&score)

		if score >= 90 && score <= 100 {
			fmt.Println("Exellent")
		} else if score >= 80 {
			fmt.Println("Good")
		} else if score >= 60 {
			fmt.Println("Not bad")
		} else {
			fmt.Println("Failed")
		}
	}


}