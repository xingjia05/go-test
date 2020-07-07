// Package algorithm ...
package algorithm

import (
	"fmt"
	"strconv"
)

// node ...
type node struct {
	id    int
	pid   int
	child []*node
}

// String ...
func (n node) String() string {
	str := ""
	str = str + "id:" + strconv.Itoa(n.id) + "pid" + strconv.Itoa(n.pid) + "\n"
	queue := []node{}
	if n.child != nil {
		for _, t := range n.child {
			queue = append(queue, *t)
		}
		for len(queue) > 0 {
			pop := queue[0]
			queue = queue[1:]
			str = str + "id:" + strconv.Itoa(pop.id) + "pid" + strconv.Itoa(pop.pid) + "\n"
			if pop.child != nil {
				for _, t := range pop.child {
					queue = append(queue, *t)
				}
			}
		}
	}
	return str
}

// ArrayToForest ...
func ArrayToForest() {
	list := initList()

	forest := node{id: -1, pid: -1}
	refer := make(map[int]*node)
	for i := 0; i < len(list); i++ {
		refer[list[i].id] = &list[i]
	}

	for i := 0; i < len(list); i++ {
		if list[i].pid == 0 {
			forest.child = append(forest.child, &list[i])
			continue
		}
		p := refer[list[i].pid]
		(*p).child = append((*p).child, &list[i])
	}
	fmt.Println(forest)
}

// init ...
func initList() []node {
	return []node{
		{id: 1, pid: 0},
		{id: 2, pid: 1},
		{id: 4, pid: 2},
		{id: 3, pid: 1}}
}
