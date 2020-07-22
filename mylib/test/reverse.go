// Package test ...
package test

import "fmt"

// listNode ...
type listNode struct {
	Next *listNode
	Val  int
}

func MyReverse() {
	list := []int{1, 2, 3, 4, 5}
	listLink := buildList(list)
	Reverse(listLink)
}

func buildList(list []int) *listNode {
	head := &listNode{Val: list[0]}
	for i := 1; i < len(list); i++ {
		item := listNode{Val: list[i]}
		item.Next = head
		head = &item
	}
	fmt.Println(head)
	fmt.Println(head.Next)
	return head
}

// Reverse ...
func Reverse(head *listNode) *listNode {
	if head == nil {
		return nil
	}
	ans := head
	c := head.Next
	ans.Next = nil
	for c != nil {
		n := c.Next
		c.Next = ans
		ans = c
		c = n
	}
	return ans
}
