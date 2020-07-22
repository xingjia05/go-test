// Package algorithm ...
package algorithm

import (
	"fmt"
)

// CanMakePaliQueries ...
func CanMakePaliQueries() {
	s := "abcda"
	queries := [][]int{{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}, {0, 4, 1}}
	ans := canMakePaliQueries(s, queries)
	fmt.Println(ans)
}

// canMakePaliQueries ...
func canMakePaliQueries(s string, queries [][]int) []bool {
	dp := make([]int, len(s))
	ans := make([]bool, len(queries))
	for i := 0; i < len(s); i++ {
		b := int(s[i]) - 97
		b ^= (1 << b)
		dp[i] = b
	}
	fmt.Println(dp)

	return ans
}
