//Package link ...
package link

import "fmt"

// MiddleNode ...
func MiddleNode() {
	a := []int{1, 2, 3, 4, 5}
	l := BuildLink(a)
	fmt.Println(l)
	middleNode(l)
}

// middleNode ...
func middleNode(head *LinkNode) *LinkNode {
	slow := head
	fast := head

	for fast.Next != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast.Next != nil {
		return slow.Next
	}
	fmt.Println(slow)
	return slow
}
