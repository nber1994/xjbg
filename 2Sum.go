package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//	res := twoSum([]int{1, 4, 6, 8}, 10)
	//	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	//	res := threeSumClosest([]int{-1, 2, 1, -4}, 1)
	//res := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	//	res := fourSum([]int{0, 0, 0, 0}, 0)
	//res := removeDuplicates([]int{1, 1, 2})
	//res := nextPermutation([]int{1, 3, 6, 4, 2})
	//res := nextPermutation([]int{7, 6, 4, 2})
	res := search([]int{1, 3}, 3)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	lens := len(nums)
	hashMap := make(map[int]int, lens)
	res := make([]int, 2)
	for k, v := range nums {
		pom := target - v
		if index, exists := hashMap[pom]; exists {
			res[0] = nums[index]
			res[1] = nums[k]
		} else {
			hashMap[v] = k
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	lens := len(nums)
	sort.Ints(nums)
	res := [][]int{}
	for left := 0; left < lens-2; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}
		mid := left + 1
		right := lens - 1
		for mid < right {
			if mid > left+1 && nums[mid] == nums[mid-1] {
				mid++
				continue
			}
			if nums[mid]+nums[left]+nums[right] == 0 {
				res = append(res, []int{nums[left], nums[mid], nums[right]})
				mid++
				right--
				continue
			}
			if nums[mid]+nums[left]+nums[right] < 0 {
				mid++
				continue
			}
			if nums[mid]+nums[left]+nums[right] > 0 {
				right--
				continue
			}
		}
	}
	return res
}

func threeSumClosest(nums []int, target int) int {
	lens := len(nums)
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for left := 0; left < lens-2; left++ {
		mid := left + 1
		right := lens - 1
		for mid < right {
			sum := nums[left] + nums[mid] + nums[right]
			if int(math.Abs(float64(sum-target))) < int(math.Abs(float64(res-target))) {
				res = sum
			}
			if sum < target {
				mid++
				continue
			} else {
				right--
				continue
			}
		}
	}
	return res
}

func fourSum(nums []int, target int) [][]int {
	lens := len(nums)
	sort.Ints(nums)
	res := [][]int{}
	for left1 := 0; left1 < lens-3; left1++ {
		if left1 > 0 && nums[left1] == nums[left1-1] {
			continue
		}
		for left2 := left1 + 1; left2 < lens-2; left2++ {
			if left2 > left1+1 && nums[left2] == nums[left2-1] {
				continue
			}
			mid := left2 + 1
			right := lens - 1
			for mid < right {
				if nums[left1]+nums[left2]+nums[mid]+nums[right] == target {
					res = append(res, []int{nums[left1], nums[left2], nums[mid], nums[right]})
					mid++
					for nums[mid] == nums[mid-1] && mid < right {
						mid++
					}
					right--
					for nums[right] == nums[right+1] && mid < right {
						right--
					}
					continue
				}
				if nums[left1]+nums[left2]+nums[mid]+nums[right] < target {
					mid++
					continue
				}
				if nums[left1]+nums[left2]+nums[mid]+nums[right] > target {
					right--
					continue
				}
			}
		}
	}
	return res
}

func removeDuplicates(nums []int) int {
	lens := len(nums)
	var i, j int
	if lens < 2 {
		return 1
	}
	for i, j = 0, 1; j < lens; j++ {
		if nums[i] == nums[j] {
			continue
		}
		i++
		nums[i] = nums[j]
	}
	return i + 1
}

//13642
func nextPermutation1(nums []int) {
	lens := len(nums)
	if lens == 0 {
		return
	}
	cut := 0
	for i := lens - 1; i > 1; i-- {
		if nums[i] > nums[i-1] {
			cut = i
			break
		}
	}
	if cut == 0 {
		for i := 0; i < lens; i++ {
			nums[i], nums[lens-1-i] = nums[lens-1-i], nums[i]
		}
	} else {
		chg1 := 0
		for i := lens - 1; i > cut-1; i++ {
			if nums[i] > nums[cut-1] {
				chg1 = i
				break
			}
		}
		nums[chg1], nums[cut-1] = nums[cut-1], nums[chg1]
		for i := lens - 1; i > (lens-1+cut)/2; i++ {
			nums[i], nums[lens-1+cut-i] = nums[lens-1+cut-i], nums[i]
		}
	}
	fmt.Println(nums)
}

func nextPermutation(nums []int) []int {
	lens := len(nums)
	if lens == 0 {
		return []int{}
	}
	index := 0
	for i := lens - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			index = i
			break
		}
	}
	if index != 0 {
		chg := index - 1
		for i := lens - 1; i > index-1; i-- {
			if nums[i] > nums[index-1] {
				chg = i
				break
			}
		}
		nums[chg], nums[index-1] = nums[index-1], nums[chg]
	}
	for i := index; i < ((lens-1+index)/2)+1; i++ {
		nums[i], nums[lens-1+index-i] = nums[lens-1+index-i], nums[i]
	}
	return nums
}

func search1(nums []int, target int) int {
	lens := len(nums)
	for left, right := 0, lens-1; left <= right; {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		if nums[mid] > nums[0] && nums[0] > target {
			left = mid + 1
		} else if target > nums[0] && nums[0] > nums[mid] {
			right = mid - 1
		} else if target > nums[mid] && nums[mid] > nums[0] {
			left = mid + 1
		} else if nums[mid] > target && target > nums[0] {
			right = mid - 1
		} else if nums[0] >= nums[mid] && nums[mid] > target {
			right = mid - 1
		} else if nums[0] > target && target > nums[mid] {
			left = mid + 1
		}
		fmt.Println(left, mid, right)
	}
	return -1
}

func search(nums []int, target int) int {
	lens := len(nums)
	left, right := 0, lens-1
	for left <= right {
		mid := (left + right) / 2
		fmt.Println(left, mid, right)
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[left] {
			if target < nums[mid] && nums[left] <= target {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
