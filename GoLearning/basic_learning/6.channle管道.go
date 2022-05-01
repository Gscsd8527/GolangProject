package main

import "fmt"

func main() {
	//1.创建channel
	ch := make(chan int, 3)

	// 2.给管道里面存储数据
	ch <- 30
	ch <- 31
	ch <- 32

	//3. 获取管道里面的数据
	a := <-ch

	fmt.Println("a = ", a)
}
