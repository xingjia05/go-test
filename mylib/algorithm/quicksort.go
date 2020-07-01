// Package algorithm ...
package algorithm

import "fmt"

// MyQK ...
func MyQK() {
	nums := []int{72, 6, 57, 88, 60, 42, 83, 73, 48, 85}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

// quickSort ...
func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	flag := nums[left]
	i := left
	j := right
	for i < j {
		for i < j && nums[j] >= flag {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= flag {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = flag
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
	return
}
