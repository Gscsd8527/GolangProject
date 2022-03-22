package main

import "fmt"

/*
	知识点一： defer的执行顺序
      当主函数退出时，defer后的函数才被调用
      defer是一个在函数结束后执行的方法，多个defer相当于压栈，先进后出

    知识点二： defer和return谁前谁后
       return先执行，defer后执行
*/



func func1()  {
	fmt.Println("A")
}

func func2()  {
	fmt.Println("B")
}

func func3()  {
	fmt.Println("c")
}

func deferFunc() int{
	fmt.Println("deferFunc....")
	return 0
}

func returnFunc() int{
	fmt.Println("returnFunc....")
	return 0
}

func returnAndDefer() int{
	fmt.Println("returnAndDefer...")
	defer deferFunc()
	return returnFunc()
}


func main() {
	//defer func1()
	//defer func2()
	//defer func3()

	fmt.Println("hello world!")
	fmt.Println("hello world!")

	returnAndDefer()
}