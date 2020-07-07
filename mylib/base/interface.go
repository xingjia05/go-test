// Paakage base ...
package base

import "fmt"

type animal interface {
	speak() string
}

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

func (cat Cat) String() string {
	return "cat speak:" + cat.speak()
}

func MyInterface() {
	cat := Cat{}
	dog := dog{}
	animal := []animal{&cat, &dog}
	fmt.Println(animal)
}
