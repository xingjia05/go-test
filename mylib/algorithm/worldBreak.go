// Package algorithm ...
package algorithm

import "fmt"

// MyWordBreak ...
func MyWordBreak() {
	str := "appleisanapple"

	wordDict := []string{"is", "apple", "an"}

	r := execute(str, wordDict)
	fmt.Println(r)
}

// execute ...
func execute(s string, wordDict []string) bool {
	sLen := len(s)

	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	dp := make(map[int]bool)
	dp[0] = true
	for i := 1; i <= sLen; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[sLen]
}
