// Package algorithm ...
package algorithm

import "fmt"

// MyIsMatch ...
func MyIsMatch() {
	//s := "aaaa"
	//p := "*a"
	//	s := "acdcab"
	//	p := "a*c?b"
	s := "acdcb"
	p := "a*c?b"

	a := isMatch(s, p)
	fmt.Println(a)
}

// isMatch ...
func isMatch(s, p string) bool {
	sLen := len(s)
	pLen := len(p)

	i := 0
	j := 0
	lasti := 0
	lastj := 0
	for {
		if i == sLen && j == pLen {
			break
		}
		if j == pLen {
			if string(p[lastj]) == "*" {
				j = lastj
				i = lasti + 1
			} else {
				return false
			}
		}
		if i == sLen {
			if string(p[j]) != "*" {
				return false
			}
			return true
		}
		if s[i] == p[j] || string(p[j]) == "?" {
			i++
			j++
			continue
		}
		if string(p[j]) == "*" {
			lasti = i
			lastj = j
			j++
			if j == pLen-1 {
				return true
			}
			continue
		}
		if p[j] != s[i] {
			if string(p[lastj]) == "*" {
				i = lasti + 1
				j = lastj
			} else {
				return false
			}
		}
	}

	return true
}
