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

func (c *Cat) speak() string {
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

<<<<<<< HEAD
=======
<<<<<<< HEAD
>>>>>>> b8197404b409a939887bac046489d1168e657f7c
func (cat Cat) String1() string {
	return "cat speak:" + cat.speak()
}

<<<<<<< HEAD
=======
=======
>>>>>>> b8197404b409a939887bac046489d1168e657f7c
// MyInterface ...
>>>>>>> 25718dc8fc7e7a0e1d2accf415f9a3b47a4b7b43
func MyInterface() {
	cat := Cat{Age: 1}
	//dog := dog{}
	pCat := &cat
	fmt.Println(pCat.Age)
	//animal := []animal{&cat, &dog}
	//fmt.Println(animal)
}
