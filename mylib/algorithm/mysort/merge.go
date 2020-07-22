// Package mysort ...
package mysort

import "fmt"

// MergeSort ...
func MergeSort() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(a[len(a):])
	r := mergeSort(a)
	fmt.Println(r)
}

// MergeSort ...
func mergeSort(r []int) []int {
	if len(r) <= 1 {
		return r
	}
	l := len(r) / 2
	left := mergeSort(r[:l])
	right := mergeSort(r[l:])
	return merge(left, right)
}

// merge ...
func merge(left, right []int) []int {
	fmt.Println(left, "--", right)
	l, r := 0, 0
	result := []int{}
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
