// Package algorithm ...
package algorithm

import "fmt"

//MinimumTotal ...
func MinimumTotal() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}

	mininumTotal(triangle)
}

// mininumTotal ...
func mininumTotal(triangle [][]int) int {
	ret := 10000
	dp := make([]int, len(triangle[len(triangle)-1]))
	dp[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := i; j >= 0; j-- {
			if j == 0 {
				dp[j] = dp[j] + triangle[i][j]
			} else if j == i {
				dp[j] = dp[j-1] + triangle[i][j]
			} else {
				dp[j] = minInt(dp[j], dp[j-1]) + triangle[i][j]
			}
		}
	}
	fmt.Println(dp)
	for _, v := range dp {
		if ret > v {
			ret = v
		}
	}
	fmt.Println(ret)
	return ret
}

// minInt ...
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
