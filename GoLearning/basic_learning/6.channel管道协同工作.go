package main

import (
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func write_channel(ch chan int){
	for i := 1; i<=10; i++{
		ch <- i
		fmt.Printf("【写入】 数据 %v 成功\n", i)
		time.Sleep(time.Millisecond * 100)
	}
	//后面有遍历，所以得关闭管道
	close(ch)
	wg2.Done()
}


func read_channel(ch chan int){
	for v:= range ch{
		fmt.Printf("【读取】 数据 %v 成功\n", v)
		time.Sleep(time.Millisecond * 50)
	}
	wg2.Done()
}


func main() {
	// 管道是安全的，读快写慢也可以，读的会等待

	var ch = make(chan int, 10)
	wg2.Add(1)
	go write_channel(ch)
	wg2.Add(1)
	go read_channel(ch)
	wg2.Wait()
	fmt.Println("结束")
}
