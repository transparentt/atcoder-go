package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)

	var b int
	fmt.Scan(&b)

	var c int
	fmt.Scan(&c)

	var x int
	fmt.Scan(&x)

	result := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				sum := 500*i + 100*j + 50*k
				if sum == x {
					result += 1
				}

			}
		}
	}

	fmt.Println(result)
}
