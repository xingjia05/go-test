// Package base ...
package base

import "fmt"

type animal interface {
	speak() string
}

// Cat ...
type Cat struct {
	Age int
}

func (c Cat) speak() string {
	return "m m"
}

type dog struct {
	name string
}

func (dog dog) speak() string {
	return "w w"
}

func (dog dog) String() string {
	return "dog speak:" + dog.speak()
}

// MyInterface ...
func MyInterface() {
	cat := Cat{}
	dog := dog{}
	animal := []animal{&cat, &dog}
	fmt.Println(animal)
}
