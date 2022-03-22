package main

import "fmt"

/*
切片的追加：
	1. 切片的长度和容量不同，长度表示左指针至右指针之间的距离，容量表示左指针至底层数组末尾的距离
	2. 切片的扩容机制，append的时候，如果长度增加后超过容量，则将容量增加两倍
切片的截取：
    跟Python中语法一致
*/

func sliceAppend(){
	//len = 3  cap = 5 slice = [0 0 0]
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d  cap = %d slice = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素1， len = 4  cap = 5 slice = [0 0 0 1]
	numbers = append(numbers, 1)
	fmt.Printf("len = %d  cap = %d slice = %v\n", len(numbers), cap(numbers), numbers)

	//len = 5  cap = 5 slice = [0 0 0 1 2]
	numbers = append(numbers, 2)
	fmt.Printf("len = %d  cap = %d slice = %v\n", len(numbers), cap(numbers), numbers)

	//len = 6  cap = 10 slice = [0 0 0 1 2 3]
	numbers = append(numbers, 3)
	fmt.Printf("len = %d  cap = %d slice = %v\n", len(numbers), cap(numbers), numbers)

}


func sliceSlice(){
	s := []int{1, 2, 3}
	s1 := s[0:2] //浅拷贝，地址没变
	fmt.Println(s1)
	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1)

	//copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3)

	// 将s中的值 依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2)
}

func main() {
	//sliceAppend()
	sliceSlice()

}
