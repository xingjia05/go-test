//Package base ...
package base

import (
	"fmt"
	"sync"
)

// Concurrence ...
func Concurrence() {
	myWG()
}

// myWG ...
func myWG() {
	wg := sync.WaitGroup{}
	i := 1
	wg.Add(1)
	go func(j int) {
		defer wg.Done()
		fmt.Println("dd", j)
	}(i)
	wg.Add(1)
	go test(i, &wg)
	wg.Wait()
}

// test ...
func test(j int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("rr", j)
}
