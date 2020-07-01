// Package algorithm ...
package algorithm

import "fmt"

// VisitOrder ...
func VisitOrder() {
	//points := [][]int{{1, 3}, {2, 4}, {3, 3}, {2, 1}}
	points := [][]int{{3, 5}, {4, 5}, {9, 1}, {5, 6}, {4, 2}}
	//direction := "LR"
	direction := "RLR"
	r := visitOrder1(points, direction)
	fmt.Println(r)
}

// VisitOrder ...
func visitOrder1(points [][]int, direction string) []int {
	pLen := len(points)
	dLen := len(direction)
	ret := []int{}

	mapP := make(map[int]bool)
	minX := points[0][0]
	minIndex := 0
	for i := 1; i < pLen; i++ {
		if points[i][0] < minX {
			minIndex = i
			minX = points[i][0]
		}
	}
	mapP[minIndex] = true
	ret = append(ret, minIndex)
	start := minIndex

	for i := 0; i < dLen; i++ {
		next := -1
		if string(direction[i]) == "L" {
			for j := 0; j < pLen; j++ {
				if mapP[j] == true {
					continue
				}
				fmt.Println(next, j, start)
				if next == -1 || cross(sub(points[next], points[start]), sub(points[j], points[start])) < 0 {
					next = j
				}
			}
			mapP[next] = true
			ret = append(ret, next)
		} else {
			for j := 0; j < pLen; j++ {
				if mapP[j] == true {
					continue
				}
				fmt.Println(next, j, start)
				if next == -1 || cross(sub(points[next], points[start]), sub(points[j], points[start])) > 0 {
					next = j
				}
			}
			mapP[next] = true
			ret = append(ret, next)
		}
		start = next
	}
	for k := 0; k < pLen; k++ {
		if mapP[k] == true {
			continue
		}
		ret = append(ret, k)
	}
	return ret
}

// sub ...
func sub(a, b []int) []int {
	return []int{a[0] - b[0], a[1] - b[1]}
}

// cross ...
func cross(a, b []int) int {
	fmt.Println(a, b, a[0]*b[1]-a[1]*b[0])
	return a[0]*b[1] - a[1]*b[0]
}

// VisitOrder ...
func visitOrder(points [][]int, direction string) [][]int {
	pLen := len(points)
	dLen := len(direction)
	ret := [][]int{}

	mapP := make(map[int]bool)
	p := points[0]
	minX := points[0][0]
	minIndex := 9999999
	for i := 1; i < pLen; i++ {
		if points[i][0] < minX {
			p = points[i]
			minIndex = i
			minX = points[i][0]
		}
	}
	mapP[minIndex] = true
	ret = append(ret, p)

	for i := 0; i < dLen; i++ {
		if string(direction[i]) == "L" {
			minY := 9999999
			for j := 0; j < pLen; j++ {
				if mapP[j] == true {
					continue
				}
				if points[j][1] < minY {
					minY = points[j][1]
					p = points[j]
					minIndex = j
				}
			}
			fmt.Println("1111")
			mapP[minIndex] = true
			ret = append(ret, p)
		} else {
			maxIndex := -9999999
			maxY := -9999999
			for j := 0; j < pLen; j++ {
				if mapP[j] == true {
					continue
				}
				if points[j][1] > maxY {
					maxY = points[j][1]
					p = points[j]
					maxIndex = j
				}
			}
			fmt.Println("2222")
			mapP[maxIndex] = true
			ret = append(ret, p)
		}
	}
	return ret
}
