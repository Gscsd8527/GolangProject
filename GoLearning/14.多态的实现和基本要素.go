package main

import "fmt"

/*
多态的基本要素
   有一个父类（有接口）
   有子类（实现了父类的全部接口方法）
   父类类型的变量（指针）指向（引用）子类的具体数据变量
*/


// 本质就是一个指针
type  AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的种类
}

//具体的类
type Cat struct{
	color string  //猫的颜色
}

func (this *Cat) Sleep(){
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string{
	fmt.Println("Cat is GetColor")
	return this.color
}

func (this *Cat) GetType() string{
	fmt.Println("Cat is GetType")
	return "Cat"
}

type Dog struct{
	color string  //猫的颜色
}

func (this *Dog) Sleep(){
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string{
	fmt.Println("Dog is GetColor")
	return this.color
}

func (this *Dog) GetType() string{
	fmt.Println("Dog is GetType")
	return "Dog"
}

func showAnimal(animalIF AnimalIF){
	animalIF.Sleep()  //多态
	fmt.Println("color = ", animalIF.GetColor())
	fmt.Println("type = ", animalIF.GetType())
}

func main() {
	//var animal AnimalIF //接口的数据类型，父类指针
	//animal = &Cat{"Red"}
	//animal.Sleep()
	//
	//animal = &Dog{"Black"}
	//animal.Sleep()


	//多态
	cat := Cat{"Yellor"}
	dog := Dog{"Write"}
	showAnimal(&cat)
	showAnimal(&dog)



}
