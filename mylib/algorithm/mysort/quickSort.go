package mysort

import "fmt"

//QuickSort ...
func QuickSort() {
	//a := []int{5, 2, 8, 4, 6, 2}
	a := []int{5, 4, 3, 2, 1}
	b := quickSort(a, 0, len(a)-1)
	fmt.Println(b)
}

func quickSort(nums []int, s int, e int) []int {
	if s >= e {
		return nums
	}
	flag := nums[s]

	left := s
	right := e

	for left < right {
		for nums[right] >= flag && left < right {
			right--
		}
		nums[left] = nums[right]
		for nums[left] <= flag && left < right {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = flag
	fmt.Println(nums, left)
	nums = quickSort(nums, s, left-1)
	nums = quickSort(nums, left+1, e)
	return nums
}
