package main

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup

func test3(n int){
	fmt.Println("NNNNN = ", n)
	for num:=(n-1)*3000 + 1; num < n*3000; num++{
		var flag = true
		for i:=2; i<num; i++{
			if num%i==0{
				flag=false
				break
			}
		}
		if flag{
			fmt.Printf("%d  是素数\n", num)
		}
	}
	wg1.Done()
}


func main() {
	for i:=1; i<4; i++{
		fmt.Println("AAAAAAaa")
		wg1.Add(1)
		go test3(i)

	}
	wg1.Wait()
}
