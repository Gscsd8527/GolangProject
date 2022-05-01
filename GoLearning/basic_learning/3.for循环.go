package main

import "fmt"

func main() {
	/*
		for 初始条件；表达条件；结束语句{
			循环体语句
		}
	*/

	// 1. 打印1-10的所有数据
	for i := 1; i <= 10; i++{
		fmt.Println("i = ", i)
	}

	// 2. for的初始语句可以被忽略，但初始语句的分号必须写
	i := 1
	for ; i<=10; i++{
		fmt.Println("i = ", i)
	}

	// 3. 注意死循环
	j := 1
	for j<= 10{
		fmt.Println("i = ", j)
		j ++
	}

	// 4. for无限循环
	//for{
	//	循环体
	//}
}
