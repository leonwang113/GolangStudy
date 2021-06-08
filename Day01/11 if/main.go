package main

import "fmt"

func main() {
	//单分支语句
	/*var is_login =true
	if is_login{
		fmt.Println("Login Success")
	}
	fmt.Println("end!")
	*/
	//双分支语句

	var is_login = false
	if is_login{
		//bool为true的时候执行的语句
		fmt.Println("login success")
	}else{
		//bool为false的时候执行的语句
		fmt.Println("login failed")
	}
	fmt.Println("end!")

	//多分支语句

	var score = 92
	if score >= 90 {
		fmt.Println("Exellent")
	}else if score >= 80 {
		fmt.Println("good")
	}else if score >= 60{
		fmt.Println("not bad")
	}else {
		fmt.Println("failed")
	}

}
