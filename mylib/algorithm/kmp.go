// Package algorithm ...
package algorithm

import "fmt"

// MyKmp ...
func MyKmp() {
	str := "132893weababaaababaaababaaababaa"
	pattern := "ababaaababaa"

	//pos := myBF(str, pattern)

	pos := myKmp(str, pattern)
	fmt.Println(pos)
}

// myKmp ...
func myKmp(s, p string) int {
	next := getNext(p)
	i := 0
	j := 0
	pos := 0
	for {
		if j == len(p) {
			pos = i - j
		}
		if i >= len(s) || j >= len(p) {
			break
		}
		if s[i] != p[j] {
			i = i - next[j] + 1
			j = next[j]
			continue
		}
		i++
		j++
	}
	return pos
}

// getNext ...
func getNext(s string) []int {
	next := make([]int, len(s)+1)
	next[0] = -1

	i := 0
	j := -1
	for {
		if i == len(s) {
			break
		}
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
			continue
		}
		j = next[j]
	}
	for k := 0; k < len(s); k++ {
		next[k]++
	}
	return next
}

// myBF ...
func myBF(s, p string) int {
	sLen := len(s)
	pLen := len(p)

	i := 0
	j := 0

	pos := 0
	for {
		if j == pLen {
			pos = i - j
			break
		}
		if i == sLen {
			break
		}
		if s[i] != p[j] {
			i = i - j + 1
			j = 0
			continue
		}
		i++
		j++
	}
	return pos
}
