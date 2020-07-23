// Package mystring ...
package mystring

import "fmt"

// BF ...
func BF() bool {
	s := "BBC ABCDAB ABCDABCDABDE"
	p := "ABCDABD"

	i := 0
	j := 0

	for {
		fmt.Println(i, j)
		if j == len(p) {
			return true
		}
		if i == len(s) {
			return false
		}
		if s[i] == p[j] {
			i++
			j++
			continue
		}
		i = i - j + 1
		j = 0
	}
}
