// Package base ...
package base

import (
	"encoding/json"
	"fmt"
)

// MyJSON ...
func MyJSON() {
	var slice []int
	jsonBytes, _ := json.Marshal(slice)
	fmt.Println(string(jsonBytes))
	slice1 := []int{}
	jsonBytes1, _ := json.Marshal(slice1)
	fmt.Println(string(jsonBytes1))
}
