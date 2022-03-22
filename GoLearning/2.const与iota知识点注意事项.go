package main

import "fmt"

//const 来定义枚举类型  iota只能在const里面使用
const(
	//可以在const() 添加一个关键词iota, 每行的iota都会累加1，第一行的iota的默认值是0
	BEIJING = 10 * iota  // tota = 0
	SHANGHAI       // tota = 1  如果第一行写了，后面不写复制操作，默认用前面的值
	SHENZHEN      // tota = 2
)

func main() {
	//常量（只读属性）
	const length int = 10
	fmt.Println("length = ", length)

	//length = 100 //常量是不允许被修改的

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)
}