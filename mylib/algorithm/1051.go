// Package algorithm ...
package algorithm

import "fmt"

// HeightChecker ...
func HeightChecker() {
	//a := []int{1, 1, 4, 2, 1, 3}
	a := []int{1, 1, 4, 2, 1, 3}
	r := heightChecker(a)
	fmt.Println(r)
}

// heightChecker 使用桶排序
func heightChecker(heights []int) int {
	bucket := make([]int, 100)

	for _, v := range heights {
		bucket[v]++
	}
	fmt.Println(bucket)
	ans := 0
	j := 0
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			if i != heights[j] {
				fmt.Println(j)
				ans++
			}
			j++
			bucket[i]--
		}
	}
	return ans
}
