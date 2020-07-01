// Package algorithm ...
package algorithm

import "fmt"

// FirstMissingPositive ...
func FirstMissingPositive() int {
	nums := []int{3, 4, -1, 1}
	fmt.Println(nums[3])

	nLen := len(nums)
	fmt.Println(nLen)
	for i := 0; i < nLen; i++ {
		fmt.Println(i)
		fmt.Println(nums[i])
		if nums[i] > 0 && nums[i] < nLen {
			nums[nums[i]] = -1
		}
	}
	fmt.Println(nums)
	for i := 1; i < nLen; i++ {
		if nums[i] > 0 {
			return i
		}
	}
	return nLen
}
