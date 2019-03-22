package main

import (
	"fmt"
)

func main() {
	a := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(a))
}

func lengthOfLIS(nums []int) int {
	lens := len(nums)
	dp := make([]int, lens)
	res := 0
	for i := 0; i < lens; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res + 1
}
