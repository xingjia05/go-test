//Package link ...
package link

import "fmt"

// ReverseList ...
func ReverseList() {
	a := []int{1, 2, 3, 4, 5}
	list := BuildLink(a)
	fmt.Println(a)
	b := reverseList(list)
	fmt.Println(b)
}

// reverseList ...
func reverseList(head *LinkNode) *LinkNode {
	ans := head
	if head == nil {
		return ans
	}
	n := head.Next
	ans.Next = nil
	for n != nil {
		t := n.Next
		n.Next = ans
		ans = n
		n = t
	}
	return ans
}
