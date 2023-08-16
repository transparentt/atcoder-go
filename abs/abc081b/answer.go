package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	integers := []int{}
	for _, s := range inputs {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		integers = append(integers, i)
	}

	result := 0
	all := true
	for all {
		for k := range integers {
			if integers[k]%2 != 0 {
				all = false
			} else {
				integers[k] = integers[k] / 2
			}
		}
		if all {
			result += 1
		}

	}

	fmt.Println(result)

}
