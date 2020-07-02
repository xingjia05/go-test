// Package mysort ...
package mysort

import "fmt"

// MyShell ...
func MyShell() {
	a := []int{1, 2, 9, 7, 4, 3, 3, 8}
	aLen := len(a)
	for g := aLen / 2; g > 0; g = g / 2 {
		for i := g; i < aLen; i++ {
			for j := i - g; j >= 0 && a[j] > a[j+g]; j = j - g {
				t := a[j]
				a[j] = a[j+g]
				a[j+g] = t
			}
		}
	}
	fmt.Println(a)
}
