package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func mergeArray() {

	n := mid
	m = last
	i = left
	j = right
	k:=0
	for i <= n && j <= m{
		if nums[i] < nums[j] {
			tmp[k++] = nums[i++]
		} else {
			tmp[k++] = nums[j++]
		}
	}
}

func mergeSort(nums []int, left, right int, tmp []int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(a, left, mid, tmp)
		mergeSort(a, mid+1, right, tmp)
		mergeArray(a, left, mid, right, tmp)
	}

}

func sort(nums []int) {
	lens = len(nums)
	tmp := []int{}
	mergeSort(nums, 0, lens-1, tmp)

}
