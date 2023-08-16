package main

import "fmt"

func main() {
	var n, y int
	fmt.Scanf("%d %d", &n, &y)

	var a, b, c int
	proof := false
	for i := 0; i < n+1; i++ {
		for j := 0; j < n-i+1; j++ {
			k := n - i - j
			sum := 10000*i + 5000*j + 1000*k
			if sum == y && k >= 0 {
				a = i
				b = j
				c = k
				proof = true
			}
		}
	}

	if proof {
		fmt.Println(a, b, c)
	} else {
		fmt.Println(-1, -1, -1)
	}
}
