// Package mysort ...
package mysort

import "fmt"

// MyHeap ...
func MyHeap() {
	a := []int{1, 2, 3, 4, 2, 1, 6, 9, 4}
	aLen := len(a)
	for i := aLen; i > 0; i-- {
		heap(a, i)
	}
	fmt.Println(a)
}

// heap ...
func heap(a []int, len int) {
	i := len/2 - 1
	var big int
	for i >= 0 {
		if i*2+1 < len && i*2+2 < len {
			big = max(a, i*2+1, i*2+2)
		} else {
			big = i*2 + 1
		}
		if a[big] > a[i] {
			t := a[big]
			a[big] = a[i]
			a[i] = t
		}
		i--
	}
	t := a[0]
	a[0] = a[len-1]
	a[len-1] = t
}

// max ...
func max(arr []int, a, b int) int {
	if arr[a] > arr[b] {
		return a
	}
	return b
}
