package main

import "fmt"

func main() {

	/*
	给出如下选项，并根据选项进行效果展示：
	    输入1：则输出"普通攻击"
	    输入2：则输出"超级攻击"
	    输入3：则输出"使用道具"
	    输入4：则输出"逃跑"
	*/
	var choice int
	fmt.Println("Please enter your choice:")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("Normal hit")
	case 2:
		fmt.Println("Super hit")
	case 3:
		fmt.Println("Use tools")
	case 4:
		fmt.Println("Run away")
	default:
		fmt.Println("Wrong input")
	}

	//如果是范围取值，则使用if else语句更为快捷；如果是确定取值，则使用switch是更优方案

	/* switch 支持多条件
	switch{
	    case 1,2:
	    default:
	}
	*/
}
