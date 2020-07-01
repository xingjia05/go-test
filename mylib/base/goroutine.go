// Package base ...
package base

import (
	"fmt"
	"time"
)

// MyGoroutine ...
func MyGoroutine() {
	for i := 1; i < 10; i++ {
		go add(i, i+1)
	}
	time.Sleep(10 * time.Second)
}

// add ...
func add(a, b int) {
	fmt.Printf("%d+%d=%d\n", a, b, a+b)
}
