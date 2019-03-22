package main

import "fmt"

func main() {
	a := []int{1, 7, 6, 1, 1, 2, 9, 5, 9, 1, 9}
	fmt.Println(mergeSort(a))
}

func mergeSort(nums []int) []int {
	lens := len(nums)
	for step := 1; step < lens; step = step * 2 {
		for i := 0; i < lens; i = i + 2*step {
			merge(nums, i, i+step, min(lens-1, i+2*step))
		}
	}
	return nums
}

func merge(nums []int, leftIndex, midIndex, rightIndex int) {
	if midIndex > len(nums)-1 {
		return
	}
	left := nums[leftIndex:midIndex]
	right := nums[midIndex:rightIndex]
	lensL := len(left)
	lensR := len(right)
	i := 0
	j := 0
	res := []int{}
	for i < lensL || j < lensR {
		if i < lensL && j < lensR {
			if left[i] < right[j] {
				res = append(res, left[i])
				i++
			} else {
				res = append(res, right[j])
				j++
			}
		} else if i < lensL {
			res = append(res, left[i])
			i++
		} else if j < lensR {
			res = append(res, right[j])
			j++
		}
	}
	for k, v := range res {
		nums[leftIndex+k] = v
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
