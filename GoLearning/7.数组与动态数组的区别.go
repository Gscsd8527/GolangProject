package main

import (
	"fmt"
)

//指定定长的数组，传过来的类型也必须是 [4]int类型
// 固定长度的数据在传参的时候，是严格匹配数组类型的
func printArray(myArray [4] int){
	fmt.Printf("printArray")
	//值拷贝
	for index, value := range myArray{
		fmt.Println("index = ", index, "value = ", value)
	}
	// 此处修改并不会影响外部的列表
	myArray[0] = 111
}


//引用传递
func printArray1(myArray []int){
	fmt.Printf("printArray1")
	// _ 表示匿名变量
	for _, value := range myArray{
		fmt.Println("value = ", value)
	}
	//值被修改了，因为动态指针传递的是地址
	myArray[0] = 111
}


func main() {
	//固定长度的数组
	var myArray1 [10] int

	myArray2 := [10]int{1, 2, 3 ,4}

	myArray3 := [4]int{11, 22 ,33 ,44}

	for i:=0; i< len(myArray1); i++{
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2{
		fmt.Println("index = ", index, "value = ", value)
	}

	//查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)

	printArray(myArray3)
	fmt.Println(myArray3)

	fmt.Println("=================")

	//动态数组，切片 slice
	myarray := []int{1, 2, 3, 4}
	fmt.Println("myarray = ", myarray)

	//动态数组在传参上是引用传递，而且不同元素长度的动态数组他们的形参也是一致的
	printArray1(myarray)
	for _, value := range myarray{
		fmt.Println("value = ", value)
	}

}