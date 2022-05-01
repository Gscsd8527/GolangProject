package main

import "fmt"

func main() {
	/*
		Go 语言中 if 条件判断的格式如下：
		if 表达式1 {
			分支1
		}else if 表达式2{
			分支2
		}else{
			分支3
		}
	*/

	//1. 最简单的if
	age := 25
	if age > 20 {
		fmt.Println("老年人")
	}

	//2. if的另一种写法: 这时候age是个局部变量，只能在内部使用
	if age:=25; age>22{
		fmt.Println("老年人")
	}


}
