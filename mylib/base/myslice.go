// Package base ...
package base

import (
	"fmt"
	"time"
)

// MySlice ...
func MySlice() {
	s0 := []int{0, 1, 2, 3}
	for len(s0) > 0 {
		time.Sleep(time.Second)
		fmt.Println(s0)
		s0 = s0[1:]
		fmt.Println(s0)
	}
}
