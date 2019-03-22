package main

import (
	"fmt"
)

func main() {
	a := map[int]int{1: 1, 2: 5, 3: 8, 4: 9, 5: 10, 6: 17, 7: 17, 8: 20, 9: 24, 10: 30}
	b := cutUp(a, 10)
	fmt.Println(b)
}

func cutDown(price map[int]int, lens int, tmpArr map[int]int) int {
	if lens == 0 {
		return 0
	}
	if a, e := tmpArr[lens]; e {
		return a
	}
	q := -1
	for i := 1; i <= lens; i++ {
		tmp := price[i] + cutDown(price, lens-i, tmpArr)
		if tmp > q {
			q = tmp
		}
	}
	tmpArr[lens] = q
	return q
}

func cutUp(price map[int]int, lens int) map[int]int {
	res := map[int]int{}
	for i := 1; i <= lens; i++ {
		q := -1
		for j := 1; j <= i; j++ {
			tmp := price[j] + res[i-j]
			if tmp > q {
				q = tmp
			}
		}
		res[i] = q
	}
	return res
}
