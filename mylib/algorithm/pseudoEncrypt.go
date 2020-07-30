// Package algorithm ...
package algorithm

import (
	"fmt"
	"math/rand"
)

// PseudoEncrypt test
func PseudoEncrypt() {
	pseudoEncrypt8(73)
}

// pseudoEncrypt ...
func pseudoEncrypt(v int) {
	l1 := v >> 16 & 0XFFFF
	r1 := v & 0XFFFF
	var l2, r2 int
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ rand.Intn(int((float64((1366*r1+150889)%714025)/714025.0)*32767))
		l1 = l2
		r1 = r2
	}
	fmt.Println(r1)
	fmt.Println((r1<<16)+l1, v)
}

// pseudoEncrypt8 stable 8 with
func pseudoEncrypt8(v int) {
	num := 13
	mask13 := (1 << num) - 1
	l1 := v >> num
	r1 := v & mask13
	var l2, r2 int
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ (rand.Intn(int((float64((1366*r1+150889)%714025)/714025.0)*32767*32767)) & mask13)
		l1 = l2
		r1 = r2
	}
	fmt.Println((r1<<num)+l1, l1, r1, v)
}
