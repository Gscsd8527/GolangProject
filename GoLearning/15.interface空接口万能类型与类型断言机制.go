package main

import "fmt"

//interface{}是万能数据类型
func myFunc(arg interface{}){
	fmt.Println("myFunc is called....")
	fmt.Println(arg)

	//interfack()该如何区分 此时引用的底层数据类型到底是什么

	//给 interface() 提供 “断言”的机制
	value, ok := arg.(string)
	if !ok{
		fmt.Println("arg 不是 string 类型")
	}else {
		fmt.Println("arg 是 string 类型, value = ", value)
		fmt.Printf("arg 类型为 %T \n", value)
	}

}

type MyBook struct {
	auth string
}


func main() {
	book := MyBook{auth: "hello"}
	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)

}