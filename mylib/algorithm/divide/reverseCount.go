// Package divide ...
package divide

import "fmt"

var num int

// ReverseCount ...
func ReverseCount() {
	a := []int{2, 4, 3, 1, 5, 6}
	fmt.Println(a)
	n := 0
	a = reverseCount(a, &n)
	fmt.Println(a)
	fmt.Println(num)
	fmt.Println(n)
}

// reverseCount ...
func reverseCount(a []int, n *int) []int {
	if len(a) <= 1 {
		return a
	}
	m := len(a) / 2
	left := reverseCount(a[:m], n)
	right := reverseCount(a[m:], n)
	return merge(left, right, n)
}

// merge ...
func merge(left []int, right []int, n *int) []int {
	ret := []int{}
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			ret = append(ret, left[i])
			i++
		} else {
			ret = append(ret, right[j])
			j++
			num = num + len(left) - i
			*n = *n + len(left) - i
		}
	}
	ret = append(ret, left[i:]...)
	ret = append(ret, right[j:]...)
	return ret
}
