// Package algorithm ...
package algorithm

import "fmt"

// FindCircleNum ...
func FindCircleNum() {
	M := [][]int{
		{1, 0, 1},
		{0, 1, 0},
		{0, 0, 1}}
	r := findCircleNum(M)
	fmt.Println(r)
}

// findCircleNum ...
func findCircleNum(M [][]int) int {
	visited := make([]int, len(M))
	queue := []int{}
	num := 0
	for i := 0; i < len(M); i++ {
		if visited[i] == 1 {
			continue
		}
		num++
		queue = append(queue, i)
		for len(queue) >= 1 {
			v := queue[0]
			queue = queue[1:]
			visited[v] = 1
			for j := 0; j < len(M); j++ {
				if M[v][j] == 1 && visited[j] == 0 {
					queue = append(queue, j)
				}
			}
		}
	}
	return num
}
