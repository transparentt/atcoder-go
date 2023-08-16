package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	allD := []int{}
	for i := 0; i < n; i++ {
		var d int
		fmt.Scan(&d)
		allD = append(allD, d)
	}

	sort.Slice(allD, func(i, j int) bool {
		return allD[i] > allD[j]
	})

	result := 0
	current := allD[0]
	for i := 1; i < len(allD); i++ {
		if current > allD[i] {
			result += 1
			current = allD[i]
		}
	}

	fmt.Println(result + 1)
}
