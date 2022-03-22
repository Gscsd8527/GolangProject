package main

import "fmt"

func main() {
	// 定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("go routine 结束")
		fmt.Println("go routine 运行")

		c <- 666  // 将666 发送给c
	}()

	num := <-c //从c中接受数据，并赋值给num

	fmt.Println("num = ", num)
	fmt.Println("main go routine 结束...")
}
