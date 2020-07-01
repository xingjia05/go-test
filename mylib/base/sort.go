// Package base ...
package base

import (
	"fmt"
	"sort"
)

// MySort ...
func MySort() {
	a := []int{1, 3, 5, 4, 2, 8, 2, 4, 5, 12, 3, 44, 12, 3, 2}
	if !sort.IntsAreSorted(a) {
		sort.Ints(a)
	}
	fmt.Println(a)
}
