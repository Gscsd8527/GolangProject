package main

import "fmt"

//通过普通方法去交换 a,b的值
//func swap(a int, b int) (int, int){
//	var temp int
//	temp = a
//	a = b
//	b = temp
//	return a, b
//}

//通过交换指针指向的地址来交换值
func swap(pa *int, pb *int){
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}


func main() {
	var a int = 10
	var b int = 20

	fmt.Println("a = ", a, "b = ", b)

	//a, b = swap(a, b)
	//fmt.Println("a = ", a, "b = ", b)

	swap(&a, &b)
	fmt.Println("a = ", a, "b = ", b)

	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int //二级指针
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
}