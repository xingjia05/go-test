package algorithm

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

// HasPathSum ...
func HasPathSum() {
	list := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1}
	//list := []int{1, -2, -3, 1, 3, -2, 0, -1}
	tree := TreeNode{Val: 5}
	queue := []*TreeNode{}
	queue = append(queue, &tree)
	i := 1
	for len(queue) > 0 && i < len(list) {
		p := queue[0]
		queue = queue[1:]
		if list[i] != 0 {
			node := TreeNode{Val: list[i]}
			(*p).Left = &node
			queue = append(queue, &node)
		}
		i++
		if i < len(list) {
			if list[i] != 0 {
				node := TreeNode{Val: list[i]}
				(*p).Right = &node
				queue = append(queue, &node)
			}
			i++
		}
	}
	y := hasPathSum(&tree, 3)
	fmt.Println(y)
}

func (n *TreeNode) String1() string {
	return strconv.Itoa((*n).Val)
}

// hasPathSum ...
func hasPathSum(root *TreeNode, sum int) bool {
	queue := []*TreeNode{}
	dpMap := make(map[*TreeNode]int)
	queue = append(queue, root)
	dpMap[root] = (*root).Val
	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]
		if now.Left == nil && now.Right == nil {
			if dpMap[now] == sum {
				return true
			}
			continue
		}
		if now.Left != nil {
			queue = append(queue, now.Left)
			dpMap[now.Left] = dpMap[now] + (*now.Left).Val
		}
		if now.Right != nil {
			queue = append(queue, now.Right)
			dpMap[now.Right] = dpMap[now] + (*now.Right).Val
		}
	}
	return false
}
