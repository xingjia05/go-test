package exam

import (
	"fmt"
	"sync"
)

// Orange ...
type Orange struct {
	Quantity int
}

// Increase ...
func (o *Orange) Increase(n int) {
	o.Quantity += n
}

// Decrease ...
func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

// String ...
func (o *Orange) String() string {
	fmt.Println("dd")
	return fmt.Sprintf("%v", o.Quantity)
	//str := "quantity:" + strconv.Itoa(o.Quantity)
	//return str
}

// Counter ...
type Counter interface {
	Increase(n int)
	//Decrease(n int)
}

// TestInterface ...
func TestInterface() {
	var o Orange
	//o := Orange{Quantity: 3}
	o.Increase(10)
	o.Decrease(5)
	fmt.Println(&o)

	var c Counter
	c = &o
	c.Increase(8)
}

// TestWaitGroup ...
func TestWaitGroup() {
	var wg sync.WaitGroup
	ch := make(chan int, 1000)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go doSomething(wg, ch, i)
	}
	wg.Wait()
	fmt.Println("All done")

	for i := 0; i < 1000; i++ {
		fmt.Println(<-ch)
	}
	var wg1 sync.WaitGroup
	wg1.Add(1)
	fmt.Println("ss")
	wg1.Done()
	wg1.Wait()
}

// doSomething ...
func doSomething(wg *sync.WaitGroup, ch chan int, i int) {
	defer wg.Done()
	ch <- i
}
