package main

import (
	"fmt"
)

//func permute(nums []int) [][]int {
//	return permuteStart(nums, 0)
//}
//
//func permuteStart(nums []int, index int) [][]int {
//	length := len(nums)
//	res := [][]int{}
//	if index == length {
//		return res
//	}
//	for i := index; i < length; i++ {
//		nums[i], nums[index] = nums[index], nums[i]
//		nextRes := permuteStart(nums, index+1)
//		nextLength := len(nextRes)
//		if 0 == nextLength {
//			res = append(res, []int{nums[index]})
//		} else {
//			for _, items := range nextRes {
//				res = append(res, append(items, nums[index]))
//			}
//		}
//		nums[i], nums[index] = nums[index], nums[i]
//	}
//	return res
//}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	lens := len(nums)
	return permuteStart(nums, 0, lens-1)
}

func permuteStart(nums []int, cursor, end int) [][]int {
	res := [][]int{}
	if cursor == end {
		return [][]int{nums}
	} else {
		for i := cursor; i <= end; i++ {
			nums[cursor], nums[i] = nums[i], nums[cursor]
			nres := permuteStart(nums, cursor+1, end)
			res = append(res, nres...)
			fmt.Println("res", res)
			nums[cursor], nums[i] = nums[i], nums[cursor]
		}
	}
	return res
}
