// Package tree ...
package tree

import (
	"fmt"
	"strconv"
)

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree ...
func BuildTree(list []int) *TreeNode {
	fmt.Println(list)
	root := &TreeNode{Val: list[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(list) {
		p := queue[0]
		queue = queue[1:]
		if list[i] != -9999 {
			item := TreeNode{Val: list[i]}
			queue = append(queue, &item)
			p.Left = &item
		}
		i++
		if i < len(list) && list[i] != -9999 {
			item := TreeNode{Val: list[i]}
			queue = append(queue, &item)
			p.Right = &item
		}
		i++
	}
	return root
}

// String ...
func (root *TreeNode) String() string {
	str := strconv.Itoa(root.Val)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p.Left != nil {
			queue = append(queue, p.Left)
			str = str + "|" + strconv.Itoa(p.Left.Val)
		} else {
			str = str + "|" + "NULL"
		}
		if p.Right != nil {
			queue = append(queue, p.Right)
			str = str + "|" + strconv.Itoa(p.Right.Val)
		} else {
			str = str + "|" + "NULL"
		}
	}
	return str
}

// MaxDepth ...
func MaxDepth() {
	a := []int{3, 9, 20, -9999, -9999, 15, 7}
	tree := BuildTree(a)
	fmt.Println(tree)
	ans := maxDepth(tree)
	fmt.Println(ans)
}

// maxDepth ...
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + MaxInt(maxDepth(root.Left), maxDepth(root.Right))
}

// MaxInt ...
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
