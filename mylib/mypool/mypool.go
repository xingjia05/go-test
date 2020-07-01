// Package mypool ...
package mypool

import (
	"fmt"
	"math/rand"
	"time"
)

var input int

// MyPool ...
func MyPool() {
	count := 2
	var ret [2]int
	index := 0
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Println("please input a value")
		fmt.Scanln(&input)
		index++
		x := rand.Intn(index)
		fmt.Println("index:", index, "rand:", x)
		if x < count {
			ret[x] = input
		}
		fmt.Println(ret)
	}
}
