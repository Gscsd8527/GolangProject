package main

import "fmt"

func main() {
	// 1. 在默认情况下，管道是双向的
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 12

	m1 := <-ch1
	m2 := <-ch1
	fmt.Println(m1, m2)

	//只读只写管道可以用在函数传参里面，表面这个函数下对管道只能读或者写

	// 2. 管道声明为只读
	ch2 := make(chan<- int, 2)
	fmt.Println(ch2)
	// 3.只写管道
	ch3 := make(<-chan int, 2)
	fmt.Println(ch3)

}
