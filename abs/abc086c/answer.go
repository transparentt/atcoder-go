package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	ts := []int{0}
	xs := []int{0}
	ys := []int{0}

	for i := 0; i < n; i++ {
		var t, x, y int
		fmt.Scanf("%d %d %d", &t, &x, &y)
		ts = append(ts, t)
		xs = append(xs, x)
		ys = append(ys, y)
	}

	can := true
	for i := 0; i < n; i++ {
		distance := int(math.Abs(float64(xs[i+1]-xs[i])) + math.Abs(float64(ys[i+1]-ys[i])))
		dt := ts[i+1] - ts[i]

		if distance > dt {
			can = false
			break
		}
		if distance%2 != dt%2 {
			can = false
			break
		}
	}

	if can {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
