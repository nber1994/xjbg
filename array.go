package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum1([]int{2, 11, 7, 15}, 9))
	fmt.Println(twoSum2([]int{2, 11, 7, 15}, 9))
	//test := [][]int{
	//	{1, 2, 3, 4},
	//	{5, 6, 7, 8},
	//	{10, 11, 12, 13},
	//	{14, 15, 16, 17},
	//}
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

//两数之和
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
//示例:
//
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//两个指针向中间逼近
func twoSum1(nums []int, target int) []int {
	lens := len(nums)
	oldNums := make([]int, lens)
	copy(oldNums[:], nums)
	left := 0
	right := lens - 1
	sortedNums := sort.IntSlice(nums)
	sort.Stable(sortedNums)
	res := []int{}
	for left < right {
		//右边的指针要移动到已经不能移动位置
		for sortedNums[left]+sortedNums[right] > target {
			right--
		}
		if sortedNums[left]+sortedNums[right] == target {
			break
		}
		//左边的指针移动到不能移动位置
		for sortedNums[left]+sortedNums[right] < target {
			left++
		}
		if sortedNums[left]+sortedNums[right] == target {
			break
		}
	}
	for k, v := range oldNums {
		if v == sortedNums[left] || v == sortedNums[right] {
			res = append(res, k)
		}
	}
	return res
}

//空间换时间
//以 差值-> 键值 作为map
func twoSum2(nums []int, target int) []int {
	myMap := map[int]int{}
	for k, v := range nums {
		if i, exist := myMap[v]; exist {
			return []int{i, k}
		}
		myMap[target-v] = k
	}
	return []int{}
}

//搜索插入位置
//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//你可以假设数组中无重复元素。
//
//示例 1:
//
//输入: [1,3,5,6], 5
//输出: 2
//示例 2:
//
//输入: [1,3,5,6], 2
//输出: 1
//示例 3:
//
//输入: [1,3,5,6], 7
//输出: 4
//示例 4:
//
//输入: [1,3,5,6], 0
//输出: 0
func searchInsert(nums []int, target int) int {
	res := 0
	for _, v := range nums {
		if v == target || v > target {
			break
		}
		res++
	}
	return res
}

//在排序数组中查找元素的第一个和最后一个位置
//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
//你的算法时间复杂度必须是 O(log n) 级别。
//
//如果数组中不存在目标值，返回 [-1, -1]。
//
//示例 1:
//
//输入: nums = [5,7,7,8,8,10], target = 8
//输出: [3,4]
//示例 2:
//
//输入: nums = [5,7,7,8,8,10], target = 6
//输出: [-1,-1]
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	for k, v := range nums {
		if v == target {
			if res[0] == -1 {
				res[0] = k
				res[1] = k
			} else {
				res[1] = k
			}
		} else {
			if res[0] == -1 {
			} else {
				break
			}
		}
	}
	return res
}

//旋转图像
//给定一个 n × n 的二维矩阵表示一个图像。
//
//将图像顺时针旋转 90 度。
//
//说明：
//
//你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
//
//示例 1:
//
//给定 matrix =
//[
//  [1,2,3],
//  [4,5,6],
//  [7,8,9]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [7,4,1],
//  [8,5,2],
//  [9,6,3]
//]
//示例 2:
//
//给定 matrix =
//[
//  [ 5, 1, 9,11],
//  [ 2, 4, 8,10],
//  [13, 3, 6, 7],
//  [15,14,12,16]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [15,13, 2, 5],
//  [14, 3, 4, 1],
//  [12, 6, 8, 9],
//  [16, 7,10,11]
//]

func rotate(matrix [][]int) {
	//首先根据中轴线进行翻转
	lens := len(matrix)
	for i := 0; i < lens; i++ {
		for j := 0; j+i < lens; j++ {
			matrix[i+j][i], matrix[i][i+j] = matrix[i][i+j], matrix[i+j][i]
		}
	}
	//再进行左右互换
	for i := 0; i < lens; i++ {
		for j := 0; j < lens/2; j++ {
			matrix[i][j], matrix[i][lens-j-1] = matrix[i][lens-j-1], matrix[i][j]
		}
	}
}

//搜索旋转排序数组
//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
//( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//
//搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//
//你可以假设数组中不存在重复的元素。
//
//你的算法时间复杂度必须是 O(log n) 级别。
//
//示例 1:
//
//输入: nums = [4,5,6,7,0,1,2], target = 0
//输出: 4
//示例 2:
//
//输入: nums = [4,5,6,7,0,1,2], target = 3
//输出: -1
func search(nums []int, target int) int {
	lens := len(nums)
	left, right := 0, lens-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
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

//搜索旋转排序数组II
//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
//( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
//
//编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
//
//示例 1:
//
//输入: nums = [2,5,6,0,0,1,2], target = 0
//输出: true
//示例 2:
//
//输入: nums = [2,5,6,0,0,1,2], target = 3
//输出: false
//进阶:
//
//这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
//这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
func searchII(nums []int, target int) bool {
	lens := len(nums)
	left, right := 0, lens-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		} else if nums[left] < nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else if nums[left] > nums[mid] {
			if nums[mid] < target && target <= nums[right] {
				left = mid
			} else {
				right = mid - 1
			}
		} else {
			left++
		}
	}
	return false
}

//无重复字符的最长子串
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
func lengthOfLongestSubstring1(s string) int {
	sa := []byte(s)
	byteMap := map[byte]int{}
	res, count := 0, 0
	for i := 0; i < len(sa); {
		if k, exist := byteMap[sa[i]]; exist {
			if count > res {
				res = count
			}
			count = 0
			i = k + 1
			byteMap = map[byte]int{}
		} else {
			byteMap[sa[i]] = i
			count++
			i++
		}
	}
	if count > res {
		res = count
	}
	return res
}

func lengthOfLongestSubstring2(s string) int {
	sa := []byte(s)
	byteMap := map[byte]int{}
	left, right, maxLen := 0, 0, 0
	for i, v := range sa {
		if k, exist := byteMap[v]; !exist {
			byteMap[v] = i
		} else {
			if byteMap[v] >= left {
				left = k + 1
			}
			byteMap[v] = i
		}
		if (right-left)+1 > maxLen {
			maxLen = right - left + 1
		}
		right++
	}
	return maxLen
}
