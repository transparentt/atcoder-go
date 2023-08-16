package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	allA := []int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		allA = append(allA, a)
	}

	sort.Slice(allA, func(i, j int) bool {
		return allA[i] > allA[j]
	})

	result := 0
	for i, v := range allA {
		if i%2 == 0 {
			result += v
		} else {
			result -= v
		}
	}

	fmt.Println(result)
}
