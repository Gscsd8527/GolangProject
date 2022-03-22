package main

import "fmt"

// 如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	// 如果说类的属性首字母大写，表示该属性是对外能够访问的，否则的话只能够类的内部访问
	Name string
	Ad   int
	Level int
}

func (this Hero) Show(){
	fmt.Println("name", this.Name)
	fmt.Println("Ad", this.Ad)
	fmt.Println("Level", this.Level)
}

func (this Hero) GetName(){
	fmt.Println("name = ", this.Name)
}


func (this Hero) SetName(newName string){
	this.Name = newName
}

func (this *Hero) SetName1(newName string){
	this.Name = newName
}

func main() {
	// 创建一个对象
	hero := Hero{Name: "张三", Ad: 100, Level: 1}
	hero.Show()

	hero.SetName("hello")
	hero.Show()

	// 传递指针，所以值被修改了
	hero.SetName1("hello")
	hero.Show()
}