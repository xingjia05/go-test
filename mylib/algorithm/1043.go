// Package algorithm ...
package algorithm

//MaxSumAfterPartitioning ...
func MaxSumAfterPartitioning() {
	A := []int{1, 15, 7, 9, 2, 5, 10}
	K := 3
	maxSumAfterPartitioning(A, K)
}

// maxSumAfterPartitioning ...
func maxSumAfterPartitioning(A []int, K int) int {
	dp := make([]int, len(A))
	dp[0] = A[0]
	for i := 1; i <= len(A); i++ {
		maxN := -1
		for j := i - 1; j >= 0 && j >= i-K; j-- {
			if A[j] > maxN {
				maxN = A[j]
			}
			if dp[i] < dp[j]+maxN*(i-j) {
				dp[i] = dp[j] + maxN*(i-j)
			}
		}
	}
	return dp[len(A)]
}
