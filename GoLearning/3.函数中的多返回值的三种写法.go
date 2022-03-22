package main

import "fmt"

// 返回一个值
func fool(a string, b int) int{
	fmt.Println("-----------fool-----------")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

// 返回多个值, 匿名
func fool2(a string, b int) (int, int){
	fmt.Println("-----------fool2-----------")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 100, 200
}


// 返回多个值, 有形参名称的
func fool3(a string, b int) (r1 int, r2 int){
	fmt.Println("-----------fool3-----------")
	// r1 和 r2 属于foo3的形参，初始化默认值为0
	// r1 和 r2 作用域空间是 foo3 整个函数体的()空间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)
	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}


// 返回多个值, 有形参名称的, 返回值如果是同类型的可以简写
func fool4(a string, b int) (r1, r2 int){
	fmt.Println("-----------fool4-----------")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}


func main() {
	c := fool("tan", 100)
	fmt.Println("c = ", c)

	ret1, ret2 := fool2("tan", 100)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

	r11, r22 := fool3("tan", 100)
	fmt.Println("r11 = ", r11, "r22 = ", r22)

	ret1, ret2 = fool4("tan", 100)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)
}