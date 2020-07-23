// Package algorithm ...
package algorithm

import "fmt"

// MyFindPeakElement ...
func MyFindPeakElement() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	r := findPeakElement(nums)
	fmt.Println(r)
}

// findPeakElement ...
func findPeakElement(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}
