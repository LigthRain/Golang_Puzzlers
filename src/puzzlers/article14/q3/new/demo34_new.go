package main

import "fmt"

type Dog struct {
	name string
}

func (dog *Dog) setName (name string) {
	dog.name = name
}

func (dog Dog) test(aa int)  {

}

func New() Dog {
	return Dog{"aaa"}
}

func main() {
	dog := &Dog{}
	dog.test(1)
	dog.setName("aaa")
	fmt.Println(dog)
	//New().setName("aaa")
}