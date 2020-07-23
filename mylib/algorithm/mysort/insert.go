// Package mysort ...
package mysort

import "fmt"

// MyInsert 12321
func MyInsert() {
	//nums := []int{1, 2, 6, 3, 4, 5}
	nums := []int{2, 3, 1}
	for i := 1; i < len(nums); i++ {
		t := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > t {
				nums[j+1] = nums[j]
				continue
			}
			break
		}
		nums[j+1] = t
	}
	fmt.Println(nums)
}
