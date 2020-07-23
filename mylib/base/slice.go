// Package base ...
package base

import "fmt"

// MySlice ...
func MySlice11() {
	array := []int{1, 2, 3, 4}
	slice := make([]int, 2)

	n := copy(slice, array)
	fmt.Println(n, slice)

	slice1 := make([]byte, 3)
	n = copy(slice1, "uasd")
	fmt.Println(n, slice1)
}

// MySlice ...
func MySlice3() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[0:2]
	slice[0] = 1000
	slice1 := array[1:2:3]
	slice1[0] = 100
	slice1 = append(slice1, 101)
	fmt.Printf("%v\n", array)
	fmt.Printf("%v,%p,%d,%d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("%v,%p,%d,%d\n", slice1, &slice1, len(slice1), cap(slice1))
}

// MySlice ...
func MySlice2() {
	s1 := []int{0, 1, 2, 3}
	s2 := [...]int{0, 1, 2, 3}
	fmt.Println(s1)
	fmt.Println(s2)
	testSlice(s1)
	testArray(s2)
	fmt.Println(s1)
	fmt.Println(s2)
}

func testSlice(x []int) {
	for i := 0; i < len(x); i++ {
		x[i] = x[i] + 100
	}
}

func testArray(x [4]int) {
	for i := 0; i < len(x); i++ {
		x[i] = x[i] + 100
	}
}

// MySlice ...
func MySlice1() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6}
	s2 := s1[2:5]
	s3 := s1[:5]
	fmt.Println(cap(s1))
	fmt.Println(cap(s2))
	fmt.Println(cap(s3))
	s2 = append(s2, 111)
	s2[1] = s2[1] + 100
	s3[2] = s3[2] + 1000
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s1)
}
