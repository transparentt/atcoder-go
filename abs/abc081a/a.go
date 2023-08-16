package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	result := 0
	if string(s[0]) == "1" {
		result += 1
	}
	if string(s[1]) == "1" {
		result += 1
	}
	if string(s[2]) == "1" {
		result += 1
	}

	fmt.Println(result)
}
