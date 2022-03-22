package main

import "fmt"
import "reflect"


func reflectNum(arg interface{}){
	fmt.Println("type : ", reflect.TypeOf(arg))
	fmt.Println("value : ", reflect.ValueOf(arg))
}

type User struct {
	Id int
	Name string
	Age int
}

func (this *User) Call(){
	fmt.Println("user is called .")
	fmt.Printf("%v\n", this)
}

func DoFiledAndMethod(input interface{}){
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("type = ", inputType)
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("value = ", inputValue)

	//通过type获取里面的字段
	// 1. 获取interface的refect.Type,通过Type得到NumFiled,进行遍历
	// 2. 得到每个field，数据类型
	// 3. 通过filed有一个Interface()方法得到 对应的value
	for i := 0; i<inputType.NumField(); i++{
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	fmt.Println("*************************")
	//通过type获取里面的方法,调用
	fmt.Println(inputType.NumMethod()) // 好像不支持使用了
	for i := 0; i<inputType.NumMethod(); i++{
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}


func main() {
	//简单测试
	//var num float64 = 1.2345
	//reflectNum(num)
	//
	user := User{Id: 1, Name: "唐三", Age: 20}
	//user.Call()

	//复杂测试
	DoFiledAndMethod(user)


}