package main

import (
	"fmt"
)

//func qsort(nums []int, left, right int) {
//	if left < right {
//		c := partition(nums, left, right)
//
//		qsort(nums, left, c-1)
//		qsort(nums, c+1, right)
//	}
//}
//
//func partition(nums []int, left, right int) int {
//	p := nums[left]
//	tmpL, tmpR := left, right
//	fmt.Println(nums[left : right+1])
//	for {
//		for nums[tmpR] > p && tmpR > tmpL {
//			tmpR--
//		}
//		for nums[tmpL] <= p && tmpL < tmpR {
//			tmpL++
//		}
//		if tmpL >= tmpR {
//			break
//		}
//		nums[tmpL], nums[tmpR] = nums[tmpR], nums[tmpL]
//	}
//	nums[tmpL], nums[left] = nums[left], nums[tmpL]
//	return tmpL
//}

func main() {
	a := []int{3, 41, 24, 76, 11, 45, 3, 3, 64, 21, 69, 19, 36}
	qsort(a)
	fmt.Println(a)
}

func qsort(nums []int) {
	qiuckSort(nums, 0, len(nums)-1)
}

func qiuckSort(nums []int, left, right int) {
	if left < right {
		m := partition(nums, left, right)

		qiuckSort(nums, left, m-1)
		qiuckSort(nums, m+1, right)
	}
}

func partition(nums []int, left, right int) int {
	p := nums[left]
	tmpL, tmpR := left, right
	for {
		for nums[tmpR] < p && tmpL < tmpR {
			tmpR--
		}
		for nums[tmpL] >= p && tmpL < tmpR {
			tmpL++
		}
		if tmpL >= tmpR {
			break
		}
		nums[tmpL], nums[tmpR] = nums[tmpR], nums[tmpL]
	}
	nums[tmpL], nums[left] = nums[left], nums[tmpL]
	return tmpL
}
