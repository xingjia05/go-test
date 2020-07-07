// Package algorithm ...
package algorithm

import (
	"fmt"
	"strconv"
)

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	str := ""
	for l != nil {
		str = str + strconv.Itoa(l.Val)
		l = l.Next
	}
	return str
}

// MyAddTwoNumbers ...
func MyAddTwoNumbers() {
	fmt.Println("111")
	l1 := initList([]int{2, 4, 3})
	fmt.Println(l1)
	l2 := initList([]int{5, 6, 4})
	fmt.Println(l2)
	l3 := addTwoNumber(l1, l2)
	fmt.Println(l3)
}

// addTwoNumber ...
func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	var pre *ListNode
	var v1, v2, last int
	for {
		if l1 == nil && l2 == nil && last == 0 {
			break
		}
		item := ListNode{Val: 0}
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}
		item.Val = (v1 + v2 + last) % 10
		last = (v1 + v2 + last) / 10
		if pre == nil {
			pre = &item
			l3 = &item
		} else {
			pre.Next = &item
			pre = pre.Next
		}
	}
	return l3
}

// initList ...
func initList(num []int) *ListNode {
	var h *ListNode
	var pre *ListNode
	for _, v := range num {
		item := ListNode{Val: v}
		if pre == nil {
			pre = &item
			h = &item
		} else {
			pre.Next = &item
			pre = &item
		}
	}
	return h
}
