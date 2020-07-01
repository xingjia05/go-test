// Package algorithm ...
package algorithm

import (
	"fmt"
)

// SpliterArray ...
func SpliterArray() {
	//nums := []int{3, 4, 8, 4, 11, 7, 3, 5, 4, 7}
	nums := []int{3, 4, 8, 4, 11, 7, 3, 4, 7}
	//n := mySplit(nums)
	n := mySplit2(nums)
	fmt.Println(n)
}

// mySplit2 ...
func mySplit2(nums []int) int {
	nLen := len(nums)
	fmt.Println(nLen)
	m := 20
	prime := make([]int, 100)
	minPrime := make([]int, 100)
	g := make([]int, 100)
	for i := 2; i < m; i++ {
		if minPrime[i] == 0 {
			prime[0]++
			prime[prime[0]] = i
			minPrime[i] = i
		}
		for j := 1; j <= prime[0]; j++ {
			if i > m/prime[j] {
				break
			}
			minPrime[i*prime[j]] = prime[j]
			if i%prime[j] == 0 {
				break
			}
		}
	}
	for j := 1; j < prime[0]; j++ {
		g[prime[j]] = nLen
	}
	fmt.Println(nLen)
	fmt.Println(g)
	fmt.Println(prime)
	fmt.Println(minPrime)
	return 0
}

// mySplit ...
func mySplit(nums []int) int {
	fmt.Println(nums)
	nLen := len(nums)
	dp := make([]int, nLen)

	dp[0] = 1

	for i := 1; i < nLen; i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			if gcd(nums[i], nums[j]) > 1 {
				dp[i] = intMin(dp[i], dp[j])
			}
		}
	}

	fmt.Println(dp)
	return dp[len(nums)-1]
}

// gcd great common divisor
func gcd(m, n int) int {
	if m > n {
		if m%n == 0 {
			return n
		}
		return gcd(n, m%n)
	}
	if n%m == 0 {
		return m
	}
	return gcd(m, n%m)
}

// intMin min
func intMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
