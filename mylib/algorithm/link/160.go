package link

import "fmt"

// GetIntersectionNode ...
func GetIntersectionNode() {
	a := []int{4, 1, 8, 4, 5}
	headA := BuildLink(a)
	fmt.Println(headA)
	b := []int{5, 0, 1, 8, 4, 5}
	headB := BuildLink(b)
	fmt.Println(headB)
	ans := getIntersectionNode(headA, headB)
	fmt.Println(ans)
}

// getIntersectionNode ...
func getIntersectionNode(headA, headB *LinkNode) *LinkNode {
	lenA := 0
	lenB := 0
	lA := headA
	lB := headB
	for lA != nil {
		lenA++
		lA = lA.Next
	}
	for lB != nil {
		lenB++
		lB = lB.Next
	}
	lA = headA
	lB = headB
	if lenA > lenB {
		l := lenA - lenB
		for i := 0; i < l; i++ {
			lA = lA.Next
		}
	} else {
		l := lenB - lenA
		for i := 0; i < l; i++ {
			lB = lB.Next
		}
	}
	for lA != nil {
		if lA == lB {
			return lA
		}
		lA = lA.Next
		lB = lB.Next
	}
	return nil
}
