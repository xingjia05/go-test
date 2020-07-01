// Package algorithm ...
package algorithm

import "fmt"

// JieYuShui 四种方法
func JieYuShui() {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(nums)

	//ans := forceRange(nums)
	//ans := dp(nums)
	ans := ponits(nums)
	fmt.Println(ans)
}

// points ...
func ponits(height []int) int {
	ans := 0
	lNum := len(height)
	leftMax := 0
	rightMax := 0
	left := 0
	right := lNum - 1

	for left < right {
		if height[left] < height[right] {
			if leftMax < height[left] {
				leftMax = height[left]
			} else {
				ans = ans + leftMax - height[left]
			}
			left++
		} else {
			if rightMax < height[right] {
				rightMax = height[right]
			} else {
				ans = ans + rightMax - height[right]
			}
			right--
		}
	}

	return ans
}

// dp ...
func dp(nums []int) int {
	ans := 0
	lNum := len(nums)
	leftMaxArr := make([]int, lNum)
	rightMaxArr := make([]int, lNum)
	for i := 1; i < lNum; i++ {
		leftMaxArr[i] = max(nums[i-1], leftMaxArr[i-1])
	}
	for i := lNum - 2; i > 0; i-- {
		rightMaxArr[i] = max(nums[i+1], rightMaxArr[i+1])
	}
	for i := 0; i < lNum; i++ {
		hight := min(leftMaxArr[i], rightMaxArr[i])
		if hight > nums[i] {
			ans = ans + hight - nums[i]
		}
	}

	return ans
}

// forceRange ...
func forceRange(nums []int) int {
	ans := 0
	lNum := len(nums)
	for i := 1; i < lNum; i++ {
		leftMax := 0
		for j := i - 1; j >= 0; j-- {
			leftMax = max(leftMax, nums[j])
		}
		rightMax := 0
		for j := i + 1; j < lNum; j++ {
			rightMax = max(rightMax, nums[j])
		}
		hight := min(leftMax, rightMax)
		if hight > nums[i] {
			ans = ans + hight - nums[i]
		}
	}
	return ans
}

// min ...
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// max ...
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
