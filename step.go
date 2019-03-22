package main

import (
	"fmt"
)

func main() {
	fmt.Println(findStep(10))
}

func findStep(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	return findStep(n-1) + findStep(n-2)
}
