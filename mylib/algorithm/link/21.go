//Package ...
package link

import "fmt"

// MergeTwoLists ...
func MergeTwoLists() {
	l1 := BuildLink([]int{1, 2, 4})
	l2 := BuildLink([]int{1, 3, 4})
	fmt.Println(l1)
	fmt.Println(l2)
	l3 := mergeTwoLists(l1, l2)
	fmt.Println(l3)
}

// mergeTwoLists ...
func mergeTwoLists(l1 *LinkNode, l2 *LinkNode) *LinkNode {
	l3 := LinkNode{}
	p := &l3
	for l1 != nil || l2 != nil {
		if l1 == nil {
			p.Next = l2
			break
		}
		if l2 == nil {
			p.Next = l1
			break
		}
		if l1.Val > l2.Val {
			p.Next = l2
			p = p.Next
			l2 = l2.Next
		} else {
			p.Next = l1
			p = p.Next
			l1 = l1.Next
		}
	}
	fmt.Println(l3)
	return l3.Next
}
