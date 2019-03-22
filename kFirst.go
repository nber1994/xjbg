package main

import (
	"fmt"
)

func main() {
	//res := findKthLargest([]int{-1, -1}, 2)
	res := findKthLargest([]int{1, 2, 3, 4, 5, 6, 7}, 4)
	fmt.Println(res)
}

//func findKthLargest(nums []int, k int) int {
//	return quickSelect(nums, 0, len(nums)-1, k)
//}
//
//func quickSelect(nums []int, left, right, k int) int {
//	if k > 0 && k <= right-left+1 {
//		pos := partition(nums, left, right)
//		fmt.Println(pos)
//		fmt.Println(nums)
//
//		if pos-left == k-1 {
//			return nums[pos]
//		}
//		if pos-left > k-1 {
//			return quickSelect(nums, left, pos-1, k)
//		}
//		return quickSelect(nums, pos+1, right, k-(pos-left+1))
//	}
//	return -1
//}
//
//func partition(nums []int, left, right int) int {
//	p := nums[left]
//	tmpl, tmpr := left, right
//	for {
//		for nums[tmpr] < p && tmpl < tmpr {
//			tmpr--
//		}
//		for nums[tmpl] >= p && tmpl < tmpr {
//			tmpl++
//		}
//		if tmpl == tmpr {
//			break
//		}
//		nums[tmpl], nums[tmpr] = nums[tmpr], nums[tmpl]
//	}
//	nums[tmpl], nums[left] = nums[left], nums[tmpl]
//	return tmpl
//}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, left, right, k int) int {
	if k < 0 || k > right-left+1 {
		return -1
	}
	pos := partition(nums, left, right)

	if pos-left+1 == k {
		return nums[pos]
	}
	if pos-left+1 > k {
		return quickSelect(nums, left, pos-1, k)
	}
	return quickSelect(nums, pos+1, right, k-(pos-left+1))

}

func partition(nums []int, left, right int) int {
	p := nums[left]
	tmpl, tmpr := left, right
	for {
		for nums[tmpr] > p && tmpr > tmpl {
			tmpr--
		}
		for nums[tmpl] <= p && tmpl < tmpr {
			tmpl++
		}
		if tmpl == tmpr {
			break
		}
		nums[tmpl], nums[tmpr] = nums[tmpr], nums[tmpl]
	}
	nums[tmpl], nums[left] = nums[left], nums[tmpl]
	return tmpl
}
