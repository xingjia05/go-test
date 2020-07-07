package algorithm

import "fmt"

// UniquePathsWithObstacles ...
func UniquePathsWithObstacles() {
	//obstacleGrid := [][]int{
	//	{0, 0, 0},
	//	{1, 1, 0},
	//	{0, 0, 0},
	//	{0, 1, 1},
	//	{0, 0, 0}}

	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0}}
	a := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(a)
}

// uniquePathsWithObstacles ...
// dp[i][i] = dp[i-1][j]+dp[i][j-1]+dp[i][j+1]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := [][]int{}
	for i := 0; i < m; i++ {
		tmpArr := []int{}
		for j := 0; j < n; j++ {
			tmpArr = append(tmpArr, 0)
		}
		dp = append(dp, tmpArr)
	}
	dp[0][0] = 1
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] > 0 {
				dp[i][j] = 0
				continue
			}
			if i-1 >= 0 {
				dp[i][j] = dp[i][j] + dp[i-1][j]
			}
			if j-1 >= 0 {
				dp[i][j] = dp[i][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
