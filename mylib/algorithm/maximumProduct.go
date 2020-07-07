// Package algorithm ...
package algorithm

// MyMaximumProduct
func MyMaximumProduct() {
	nums := []int{1, 2, 3, 4}
	maximumProduct(nums)
}

// maximumProduct ...
func maximumProduct(nums []int) int {
	maxList := []int{-1001, -1001, -1001}
	smallList := []int{1001, 1001}
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxList[0] {
			maxList[0] = nums[i]
		}
		if maxList[0] > maxList[1] {
			maxList[0], maxList[1] = maxList[1], maxList[0]
		}
		if maxList[1] > maxList[2] {
			maxList[1], maxList[2] = maxList[2], maxList[1]
		}

		if nums[i] < smallList[0] {
			smallList[0] = nums[i]
		}
		if smallList[0] < smallList[1] {
			smallList[0], smallList[1] = smallList[1], smallList[0]
		}
	}

	if maxList[0]*maxList[1] > smallList[0]*smallList[1] {
		return maxList[0] * maxList[1] * maxList[3]
	}
	return smallList[0] * smallList[1] * maxList[3]

}
