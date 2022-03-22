package main

import "fmt"

func main() {
	// 第一种： 声明slice1是一个切片，并且初始化，默认值是1,2,3，长度len是3
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 第二种： 声明slice1是一个切片，但是并咩有给slice分配空间
	var slice2 []int
	slice2 = make([]int, 3) //开辟三个空间，默认值为0
	slice2[0] = 100
	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2)

	//第三种: 声明slice1是一个切片, 同时给slice分配空间，三个空间，初始化值为0
	var slice3 []int = make([]int, 3)
	fmt.Printf("len = %d, slice3 = %v\n", len(slice3), slice3)

	//第四种：声明slice1是一个切片, 同时给slice分配空间，三个空间，初始化值为0,通过 := 推导出slice是一个切片
	slice4 := make([]int, 3)
	fmt.Printf("len = %d, slice4 = %v\n", len(slice4), slice4)

	// 判断一个silce为0, 表示里面没有存放任何元素
	var slice5 []int
	if slice5 == nil{
		fmt.Println("slice5是一个空切片")
	}else {
		fmt.Println("slice5是有空间的")
	}
}