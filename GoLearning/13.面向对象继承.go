package main

import "fmt"

type Human struct {
	name string
	sex string
}

func (this * Human) Eat(){
	fmt.Println("Human.Eat...")
}

func (this * Human) Walk(){
	fmt.Println("Human.Walk...")
}

type SuperMan struct {
	Human //SuperMan类继承了Human类的方法
	level int
}

// 重定义父类的方法Eat()
func (this *SuperMan) Eat(){
	fmt.Println("SuperMan.Eat...")
}

// 子类的新方法
func (this *SuperMan) Fly(){
	fmt.Println("SuperMan.Fly...")
}

func main() {
	h := Human{name: "tan", sex: "nan"}
	h.Eat()
	h.Walk()

	//定义一个子类对象
	s := SuperMan{Human{"李四", "nannann"}, 100}
	s.Walk() // 父类方法
	s.Eat()  // 重写方法
	s.Fly()  // 子类方法
}