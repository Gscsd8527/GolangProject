package main

import "fmt"

/*
			变量
		type          value     (pair)
static type  concrete type

每个变量都有 类型和值
类型又分为 static type 和 concrete type

*/

func main() {
	var a string
	//pari<statictype:string, value: "abc">
	a = "abc"

	//pari<type:string, value: "abc">
	var allType interface{}
	allType = a
	fmt.Println(allType)


	value, _ := allType.(string)
	fmt.Println(value)
}
