// Package search ...
package search

import "fmt"

// Bsearch ...
func Bsearch() {
	a := []int{1, 2, 4, 5, 8, 9, 90}

	low := 0
	high := len(a)
	target := 90
	var ans int
	for low <= high && high >= 0 && low < len(a) {
		mid := (low + high) / 2
		if a[mid] == target {
			ans = mid
			break
		}
		if a[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println(ans)
}
