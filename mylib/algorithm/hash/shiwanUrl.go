// Package hash ...
package hash

import "fmt"

// ShiWanUrl ...
func ShiWanUrl() {
	sets := []string{"a", "f", "a", "f", "f", "s", "d"}
	fmt.Println(sets)

	strMap := make(map[string]int, len(sets))

	for _, v := range sets {
		strMap[v] = strMap[v] + 1
	}
	fmt.Println(strMap)
}
