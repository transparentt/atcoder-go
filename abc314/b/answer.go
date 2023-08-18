package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	var n int
	fmt.Scan(&n)

	bets_all := [][]int{}
	for i := 0; i < n; i++ {
		var c int
		fmt.Scan(&c)
		bets := []int{}
		for i := 0; i < c; i++ {
			var a int
			fmt.Scan(&a)
			bets = append(bets, a)
		}
		bets_all = append(bets_all, bets)
	}
	var x int
	fmt.Scan(&x)

	ids := []int{}
	lengths := []int{}
	for i, bets := range bets_all {
		if slices.Contains(bets, x) {
			ids = append(ids, i+1)
			lengths = append(lengths, len(bets))
		}
	}

	minValue := 150
	for _, v := range lengths {
		if minValue > v {
			minValue = v
		}
	}

	result := []string{}
	for i, v := range lengths {
		if v == minValue {
			result = append(result, fmt.Sprint(ids[i]))
		}
	}

	sort.Slice(result, func(i, j int) bool {
		k, _ := strconv.Atoi(result[i])
		l, _ := strconv.Atoi(result[j])
		return k < l
	})

	fmt.Println(len(result))
	fmt.Println(strings.Join(result, " "))

}
