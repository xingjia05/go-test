package base

import (
	"fmt"
	"strconv"
)

// node ...
type node struct {
	Val  int
	Next *node
}

func (n *node) String() string {
	str := "head"
	for n != nil {
		str = str + "->" + strconv.Itoa(n.Val)
		n = n.Next
	}
	return str
}

// MyPointer ...
func MyPointer() {
	list := []int{1, 2, 3}
	head := &node{}
	listP := []*node{}
	for i := 0; i < len(list); i++ {
		item := node{Val: list[i]}
		item.Next = head.Next
		head.Next = &item
		listP = append(listP, &item)
	}
	fmt.Println(head.Next)
	t := listP[0]
	(*t).Val = 100
	fmt.Println(t)
	fmt.Println(listP)
}
