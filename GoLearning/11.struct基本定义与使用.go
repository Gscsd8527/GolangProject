package main

import "fmt"

//声明一种新的数据类型 myint, 是int的一个别名
type myint int

//定义一个结构体
type Book struct {
	title string
	auth string
}

func changeBook(book Book){
	book.auth = "my"
}

func changeBook2(book *Book){
	book.auth = "my"
}

func main() {
	//var a myint = 10
	//println("a = ", a)

	var book Book
	book.title = "百度"
	book.auth = "李彦宏"
	fmt.Println("变更前： ",book)
	changeBook(book)
	fmt.Println("变更后： ",book)

	changeBook2(&book)
	fmt.Println("变更后： ",book)
}
