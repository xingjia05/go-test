// Package link ...
package link

import (
	"fmt"
	"strconv"
)

type LinkNode struct {
	Val  int
	Next *LinkNode
}

// Stirng ...
func (linkNode *LinkNode) String() string {
	str := ""
	for linkNode != nil {
		str = str + "->" + strconv.Itoa(linkNode.Val)
		linkNode = linkNode.Next
	}
	return str
}

// BuildLink ...
func BuildLink(a []int) *LinkNode {
	var head *LinkNode
	for i := len(a) - 1; i >= 0; i-- {
		item := LinkNode{Val: a[i]}
		if head == nil {
			head = &item
			continue
		}
		item.Next = head
		head = &item
	}
	return head
}

// HuiWen ...
func HuiWen() bool {
	a := []int{1, 2, 3, 4, 5, 6}
	list := BuildLink(a)
	fmt.Println(list)

	f := list
	t := list
	for t.Next != nil && t.Next.Next != nil {
		f = f.Next
		t = t.Next.Next
	}
	r := f.Next
	c := f.Next.Next
	r.Next = nil
	for c != nil {
		tt := c.Next
		c.Next = r
		r = c
		c = tt
	}
	fmt.Println(r)
	head := list
	for r != nil {
		if r.Val != head.Val {
			fmt.Println("false")
			return false
		}
		r = r.Next
		head = head.Next
	}
	return true
}
