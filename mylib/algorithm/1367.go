// Package algorithm ...
package algorithm

import (
	"fmt"
	"strconv"
)

//ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// String ...
func (listNode *ListNode) String() string {
	str := ""
	for listNode != nil {
		str = str + "|" + strconv.Itoa(listNode.Val)
		listNode = listNode.Next
	}
	return str
}

// String ...
func (treeNode *TreeNode) String() string {
	str := ""
	queue := []*TreeNode{}
	queue = append(queue, treeNode)
	str = str + strconv.Itoa(treeNode.Val)
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p.Left != nil {
			queue = append(queue, p.Left)
			str = str + "|" + strconv.Itoa(p.Left.Val)
		} else {
			str = str + "|" + "null"
		}
		if p.Right != nil {
			queue = append(queue, p.Right)
			str = str + "|" + strconv.Itoa(p.Right.Val)
		} else {
			str = str + "|" + "null"
		}
	}
	return str
}

// IsSubPath ...
func IsSubPath() {
	root := buildTree()
	head := buildList()
	fmt.Println(head)
	fmt.Println(root)
	ans := isSubPath(head, root)
	fmt.Println(ans)
}

// isSubPath ...当存在两个一样的分支节点的时候使用stack存储，做回归用
func isSubPath(head *ListNode, root *TreeNode) bool {
	fmt.Println("root", root)
	if root == nil {
		return false
	}
	return isSubPath1(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

// isSubPath1 ...当存在两个一样的分支节点的时候使用stack存储，做回归用
func isSubPath1(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val == root.Val {
		head = head.Next
	} else {
		return false
	}

	ans := isSubPath1(head, root.Left)
	if ans == true {
		return true
	}
	ans = isSubPath1(head, root.Right)
	if ans == true {
		return true
	}
	return false
}

// buildList ...
func buildList() *ListNode {
	//list := []int{4, 2, 8}
	//list := []int{1, 4, 2, 6, 8}
	//list := []int{1, 3}
	//list := []int{1, 10}
	list := []int{4, 2, 1}
	var head *ListNode
	for i := len(list) - 1; i >= 0; i-- {
		item := ListNode{Val: list[i]}
		if head == nil {
			head = &item
		} else {
			item.Next = head
			head = &item
		}
	}
	return head
}

// buildTree ...
func buildTree() *TreeNode {
	//list := []int{1, 4, 4, -1, 2, 2, -1, 1, -1, 6, 8, -1, -1, -1, -1, 1, 3}
	//list := []int{1, 4, 2, 3}
	//list := []int{1, -1, 1, 10, 1, 9}
	list := []int{9, -1, 4, 3, -1, 1, -1, 7, -1, 1, 2}
	root := TreeNode{Val: list[0]}
	queue := []*TreeNode{&root}
	i := 1
	for len(queue) > 0 && i < len(list) {
		p := queue[0]
		queue = queue[1:]
		if list[i] != -1 {
			item := TreeNode{Val: list[i]}
			p.Left = &item
			queue = append(queue, &item)
		}
		i++
		if i < len(list) && list[i] != -1 {
			item := TreeNode{Val: list[i]}
			p.Right = &item
			queue = append(queue, &item)
		}
		i++
	}
	return &root
}
