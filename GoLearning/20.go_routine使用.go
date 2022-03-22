package main

import (
	"fmt"
	"time"
)

func newTask(){
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}


func main() {
	// 创建一个go程 去执行newTask流程
	go newTask()

	fmt.Println("********************")

	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
