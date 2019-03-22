package main

import "fmt"

func main() {
	a := []int{2, 7, 9, 4, 1}
	fmt.Println(zhifangtu(a))
}

func zhifangtu(he []int) int {
	lens := len(he)
	dp := make([][100]int, 100)
	for i := 0; i < lens; i++ {
		dp[i][i] = he[i]
	}
	for i := 1; i < lens; i++ {
		for j := 0; j+i < lens; j++ {
			min := he[j]
			for k := j; k <= j+i; k++ {
				if he[k] < min {
					min = he[k]
				}
			}
			max := dp[j+1][i+j]
			if dp[j][i+j-1] > dp[j+1][i+j] {
				max = dp[j][i+j-1]
			}
			if max < min*(j+1) {
				max = min * (j + 1)
			}
			dp[j][i+j] = max
		}
	}
	return dp[0][lens-1]
}

func zhifangtu(nums []int) int {
	lens := len(nums)
	dp = make([][100]int, 100)
	for i := 0; i < lens; i++ {
		dp[i][i] = nums[i]
	}
	for gap := 1; gap < lens; gap++ {
		for i := 0; i+gap < lens; i++ {
			min = nums[i]
			for j := i; j <= i+gap; j++ {
				if nums[j] < min {
					min = nums[j]
				}
			}
			max := dp[i+1][i+gap]
			if dp[i+1][i+gap] < dp[i][i+gap-1] {
				max = dp[i][i+gap-1]
			}
			if max < min*(gap+1) {
				max = min * (gap + 1)
			}
			dp[i][i+gap] = max
		}
	}
	return dp[0][lens-1]
}
