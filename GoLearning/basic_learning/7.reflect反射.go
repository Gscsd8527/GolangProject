package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
}

func reflectDn(x interface{}){
	v := reflect.TypeOf(x)
	fmt.Println(v)
}

func main() {
	a := 1
	b := "2"
	c := 3.3
	d := true
	reflectDn(a)
	reflectDn(b)
	reflectDn(c)
	reflectDn(d)
	reflectDn(Person{})
}
