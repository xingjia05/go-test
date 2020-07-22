// Package base ...
package base

import "fmt"

// MyArray ...
func MyArray() {
	m := 2
	n := 3
	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	fmt.Println(arr)
}

// MyArray1 ...
func MyArray1() {
	m := 2
	n := 3
	arr := [][]int{}
	for i := 0; i < m; i++ {

		t := []int{}
		for j := 0; j < n; j++ {
			t = append(t, 10)
		}
		arr = append(arr, t)
	}
	fmt.Println(arr)
}
