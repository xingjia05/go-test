// Package link  ...
package link

import "fmt"

// RemoveNthFromEnd ...
func RemoveNthFromEnd() {
	a := []int{1, 2, 3, 4, 5}

	l := BuildLink(a)
	l1 := removeNthFromEnd(l, 2)
	fmt.Println(l1)
}

// removeNthFromEnd ...
func removeNthFromEnd(head *LinkNode, n int) *LinkNode {
	fir := head
	sec := head
	for fir != nil {
		fir = fir.Next
		if 0 > n {
			sec = sec.Next
		}
		n--
	}
	if n == 0 {
		head = head.Next
	} else if sec.Next != nil {
		sec.Next = sec.Next.Next
	} else {
		sec.Next = nil
	}
	return head
}
