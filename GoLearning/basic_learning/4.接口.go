package main

import "fmt"

type Usber interface {
	Start()
	Stop()
}


// 如果接口里面有方法的话，必须要通过结构体或者自定义类型实现这个接口

type Phone struct {
	Name string
}

func (p *Phone) Start(){
	fmt.Printf("%s  start  \n", p.Name)
}

func (p *Phone) Stop(){
	fmt.Printf("%s  stop  \n", p.Name)
}

type Computer struct {
}

func (c Computer) work(usb Usber){
	usb.Start()
	usb.Stop()
}

func main() {
	var usb Usber
	p := Phone{Name: "华为"}
	usb = &p
	usb.Start()
	usb.Stop()

	var computer = Computer{}
	var phone = Phone{
		Name: "苹果",
	}
	computer.work(&phone)

}
