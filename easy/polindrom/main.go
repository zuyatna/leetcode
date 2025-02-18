package main

import "fmt"

func main() {
	fmt.Println(isPolindrom(121))
	fmt.Println(isPolindrom(-121))
}

func isPolindrom(x int) bool {
	if x < 0 {
		return false
	}

	var reversed int
	original := x

	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return original == reversed
}
