// Package mystring ...
package mystring

import "fmt"

// KMP ...
func KMP() {
	s := "BBC ABCDAB ABCDABCDABDE"
	p := "ABCDABD"

	r := kmp(s, p)
	fmt.Println(r)
}

// kmp ...
func kmp(s string, p string) int {
	fmt.Println(s)
	fmt.Println(p)

	next := getNext(p)
	fmt.Println(next)
	i := 0
	j := 0
	for {
		fmt.Println(i, j)
		if j == len(p) {
			return i - j
		}
		if i == len(s) {
			return -1
		}
		if j == -1 || s[i] == p[j] {
			i++
			j++
			continue
		}
		j = next[j]
	}
}

// getNext ...
func getNext(p string) []int {
	ans := make([]int, len(p))
	ans[0] = -1
	i := 1
	j := 0
	for i < len(p)-1 {
		if j == -1 || p[i] == p[j] {
			j++
			i++
			ans[i] = j
			continue
		}
		j = ans[j]
	}
	return ans
}
