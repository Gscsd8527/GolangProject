package main

import "fmt"

func main() {
	c := make(chan int, 3)  //带有缓冲的channel

	fmt.Println("len(c)= ", len(c), "cap(c)= ", cap(c))

	go func() {
		defer fmt.Println("子 go程 结束")
		for i:=0; i<3; i++{
			c <- i
			fmt.Println("子go程 正在运行 len(c)= ", len(c), "cap(c)= ", cap(c))
		}
	}()
	for i:=0; i<3; i++{
		num := <-c
		fmt.Println("num = ", num)
	}
	fmt.Println("main 结束")
}
