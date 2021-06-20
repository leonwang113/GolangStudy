package main

import "fmt"

func main() {

	//case1
	data := map[string][]string{"hebei":[]string{"廊坊","石家庄","邯郸"}, "beijing":[]string{"朝阳","丰台","海淀"} }
	fmt.Println(data["beijing"][0])

	//case2
	info := map[int]map[string]string{1001:{"name":"leon","age":"23"}, 1002:{"name":"alvin","age":"33"}}
	fmt.Println(info[1002]["name"])

	//case3
	info2 := []map[string]string{{"name":"alvin1","sid":"1001"},{"name":"alvin2","sid":"1002"},{"name":"alvin3","sid":"1003"}}
	fmt.Println(info2[2]["sid"])
}
