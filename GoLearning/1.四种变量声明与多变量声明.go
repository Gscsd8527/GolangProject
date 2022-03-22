package main

import "fmt"

// 声明全局变量 方法一、方法二、方法三是可以的
//var gA int = 100
//var gB = 200

// 用方法四来声明全局变量（不允许）
//gC := 200  := 只能在函数体内使用


func main() {
	//方法一： 声明一个变量 默认值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("a 类型为 %T\n", a)

	//方法二： 声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("b 类型为 %T\n", b)

	//方法三： 在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("c 类型为 %T\n", c)

	//方法四：（常用的方法）省去 var关键词，直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("e 类型为 %T", e)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, "yy = ", yy)

	var kk, ll = 100, "tan"
	fmt.Println("kk = ", kk, "ll = ", ll)

	//多行的多变声明
	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, "jj = ", jj)

}
