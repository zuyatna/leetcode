package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fast
	fmt.Println(isPolindrom(121))
	fmt.Println(isPolindrom(-121))

	// slow
	fmt.Println(isPolindrom2(121))
	fmt.Println(isPolindrom2(-121))
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

func isPolindrom2(x int) bool {
	str := strconv.Itoa(x)

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
