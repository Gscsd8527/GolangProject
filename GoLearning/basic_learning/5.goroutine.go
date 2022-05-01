package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test1(){
	i := 1
	for{
		fmt.Println("test1 你好 golang", i)
		if i >= 10{
			break
		}
		time.Sleep(time.Millisecond * 100)
		i ++
	}
	wg.Done() //携程计数器-1
}

func test2(){
	i := 1
	for{
		fmt.Println("test2 你好 golang", i)
		if i >= 10{
			break
		}
		time.Sleep(time.Millisecond * 100)
		i ++
	}
	wg.Done() //携程计数器-1
}

func main() {
	wg.Add(1)
	go test1() //开启一个携程
	wg.Add(1)
	go test2() //开启一个携程
	wg.Wait()
	fmt.Println("主线程退出")
}