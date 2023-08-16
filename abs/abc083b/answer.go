package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	passed := []int{}
	for i := 1; i <= n; i++ {
		str := fmt.Sprint(i)

		sumStr := 0
		for j := 0; j < len(str); j++ {
			integer, err := strconv.Atoi(string(str[j]))
			if err != nil {
				panic(err)
			}
			sumStr += integer
		}

		if sumStr >= a && sumStr <= b {
			passed = append(passed, i)
		}
	}

	result := 0
	for _, p := range passed {
		result += p
	}

	fmt.Println(result)

}
