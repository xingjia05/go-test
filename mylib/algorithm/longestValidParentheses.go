// Package algorithm ...
package algorithm

import "fmt"

// MyLongestValidParentheses ...
func MyLongestValidParentheses() {
	//str := ")()())"
	str := "()(()"
	r := longestValidParentheses(str)
	fmt.Println(r)
}

// longestValidParentheses ...
func longestValidParentheses(s string) int {
	return stack(s)
	//return myDp(s)
}

// myDp ...
func myDp(s string) int {
	sLen := len(s)
	dpMap := make([]int, sLen)
	maxAns := 0
	for i := 1; i < sLen; i++ {
		if string(s[i]) == "(" {
			continue
		}
		if string(s[i-1]) == "(" {
			if i < 2 {
				dpMap[i] = 2
			} else {
				dpMap[i] = dpMap[i-2] + 2
			}
		} else if i-dpMap[i-1] > 0 && string(s[i-dpMap[i-1]-1]) == "(" {
			if i-dpMap[i-1] >= 2 {
				dpMap[i] = dpMap[i-1] + dpMap[i-dpMap[i-1]-2] + 2
			} else {
				dpMap[i] = dpMap[i-1] + 2
			}
		}
		maxAns = mymax(maxAns, dpMap[i])
	}
	return maxAns
}

// mymax ...
func mymax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// stack ...
func stack(s string) int {
	stack := []int{}
	stack = append(stack, -1)
	sLen := len(s)
	max := 0
	for i := 0; i < sLen; i++ {
		if string(s[i]) == "(" {
			stack = append(stack, i)
			continue
		}
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			max = mymax(max, i-stack[len(stack)-1])
		}
	}
	return max
}

// my ...
func my(s string) int {
	return 0
}
