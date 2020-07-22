// Package link ...
package link

import "fmt"

// HasCycle ...
func HasCycle() {
	a := []int{3, 2, 0, -4}
	pos := 1

	list := BuildLink(a)
	fmt.Println(list)

	var tail *LinkNode
	tail = list
	for tail.Next != nil {
		tail = tail.Next
	}
	var c *LinkNode
	c = list
	for i := 0; i < pos; i++ {
		c = c.Next
	}
	fmt.Println(c)
	fmt.Println(tail)
	tail.Next = c
	b := hasCycle(list)
	fmt.Println(b)
}

// hasCycle ...
func hasCycle(head *LinkNode) bool {
	l1 := head
	l2 := head

	for l1 != nil && l2.Next != nil {
		l1 = l1.Next
		l2 = l2.Next.Next
		if l1 == l2 {
			return true
		}
	}
	return false

}
