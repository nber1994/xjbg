package main

import "fmt"

type test struct {
	a []int
	b int
}

func main() {
	pre := nil
	pre = &test{[]int{}, 0}
	fmt.Println("ok")
}
