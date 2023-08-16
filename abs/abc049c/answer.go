package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	for {
		if strings.HasSuffix(s, "dream") {
			s = s[:len(s)-len("dream")]
		} else if strings.HasSuffix(s, "dreamer") {
			s = s[:len(s)-len("dreamer")]
		} else if strings.HasSuffix(s, "erase") {
			s = s[:len(s)-len("erase")]
		} else if strings.HasSuffix(s, "eraser") {
			s = s[:len(s)-len("eraser")]
		} else {
			break
		}
	}

	if len(s) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
