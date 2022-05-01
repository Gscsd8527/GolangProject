package main

import (
	"fmt"
)


// 1. 在同一个输出函数中输出多项的时候， Hello World! 和 Golang 之间是存在空格的
// 2. 不同函数之间会换行

func FuncPrintLn(){
	fmt.Println("Hello World!", "Golang")
	fmt.Println("Hello World!")
	//Hello World! Golang
	//Hello World!
}

// 1. 在同一个输出函数中输出多项的时候， Hello World! 和 Golang 之间不存在空格的
// 2. 在不同输出函数之间不换行

func FuncPrint(){
	fmt.Print("Hello World!", "Golang")
	fmt.Print("Hello World!")
	//Hello World!GolangHello World!
}


//1. 可以对参数进行格式化输出，在不同输出函数中是不换行的

func FuncPrintf(){
	a := 10
	b := 20
	c := "hello"
	fmt.Printf("a = %d, b = %d", a, b)
	fmt.Printf("c = %s", c)
	//a = 10, b = 20 c = hello
}


func main() {
	FuncPrintLn()
	FuncPrint()
	FuncPrintf()

}