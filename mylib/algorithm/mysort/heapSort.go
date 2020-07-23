// Package mysort ...
package mysort

import "fmt"

// HeapSort ...
func HeapSort() {
	a := []int{0, 7, 5, 19, 8, 4, 1, 20, 13, 16}
	fmt.Println(a)
	n := len(a) - 1
	a = buildHeap(a, n)
	fmt.Println(a)
	for i := n; i >= 1; i-- {
		swap(a, 1, i)
		heapify(a, i-1, 1)
	}
	fmt.Println(a)
}

// buildHeap
func buildHeap(a []int, n int) []int {
	for i := n / 2; i >= 1; i-- {
		heapify(a, n, i)
	}
	return a
}

// heapify ...
func heapify(a []int, n int, i int) {
	fmt.Println(n, i)
	for {
		maxpos := i
		if i*2 <= n && a[maxpos] < a[2*i] {
			maxpos = 2 * i
		}
		if i*2+1 <= n && a[maxpos] < a[2*i+1] {
			maxpos = 2*i + 1
		}
		if maxpos == i {
			break
		}
		swap(a, i, maxpos)
		i = maxpos
	}
}

// swap ...
func swap(a []int, i int, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}
