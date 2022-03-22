package main

import "fmt"

//遍历map： 引用传递
func printMap(cityMap map[string]string){
	for key, value := range cityMap{
	fmt.Println("key =", key, "value =", value)
	}
}

//修改map： 引用传递
func changeMap(cityMap map[string]string){
	cityMap["等级"] = "100"
}


func mapSm(){
	// 声明myMap是一种map类型 key是string, value是string
	var myMap1 map[string]string
	if myMap1 == nil{
	println("myMap1 是一个空map")
	}
	//在使用map前，需要先用make给map分配数据空间
	myMap1 = make(map[string]string, 10)
	myMap1["name"] = "唐三"
	myMap1["sex"] = "男"
	myMap1["age"] = "20"
	fmt.Println(myMap1)

	//第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "唐三"
	myMap2[2] = "唐三"
	myMap2[3] = "唐三"
	fmt.Println(myMap2)

	//第三种声明方式
	myMap3 := map[string]string{
	"1": "1",
	"2": "2",
	"3": "3",
	}
	fmt.Println(myMap3)
}


func mapSy(){
	cityMap := make(map[string]string)

	//添加
	cityMap["name"] = "唐三"
	cityMap["sex"] = "男"
	cityMap["age"] = "20"

	//遍历
	for key, value := range cityMap{
		fmt.Println("key =", key, "value =", value)
	}

	//删除
	delete(cityMap, "age")
	fmt.Println(cityMap)

	//修改
	cityMap["name"] = "斗罗大陆"
	fmt.Println(cityMap)
}



func main() {
	//mapSm()
	mapSy()
}
