package main

import (
	"fmt"
)

type MinHeap struct {
	heap []int
	size int
}

//func findKthLargest(nums []int, k int) int {
//	mh := &MinHeap{[]int{}, 0}
//}

func (this *MinHeap) Push(v int) {
	this.heap = append(this.heap, v)
	this.size++
	this.autoFixUp()
}

func (this *MinHeap) autoFixUp() {
	if 0 == this.size {
		return
	}
	v := this.heap[this.size-1]
	i := this.size - 1
	j := (i - 1) / 2
	for i > 0 {
		if this.heap[j] <= v {
			break
		}
		this.heap[i], this.heap[j] = this.heap[j], this.heap[i]
		i = j
		j = (i - 1) / 2
	}
	//this.heap[i] = v
}

func (this *MinHeap) Pop() int {
	if this.size == 0 {
		return -1
	}
	res := this.heap[0]
	this.size--
	this.autoFixDown()
	return res
}

func (this *MinHeap) autoFixDown() {
	if 0 == this.size {
		return
	}
	this.heap[0] = this.heap[this.size]
	this.heap = this.heap[0:this.size]
	fmt.Println(this.heap)
	i := 0
	j := 0
	for i < this.size {
		child := -1
		if i*2+2 < this.size {
			if this.heap[i*2+2] < this.heap[i*2+1] {
				child = this.heap[i*2+2]
				j = i*2 + 2
			} else {
				child = this.heap[i*2+1]
				j = i*2 + 1
			}
		} else if i*2+1 < this.size {
			child = this.heap[i*2+1]
			j = i*2 + 1
		}
		if child > 0 && child < this.heap[i] {
			this.heap[j], this.heap[i] = this.heap[i], this.heap[j]
			i = j
		} else {
			break
		}
	}
}

func main() {
	a := []int{5, 4, 6, 7, 1, 3}
	a = []int{5, 7, 3, 1, 9, 6, 2, 4, 8, 0}
	mh := &MinHeap{[]int{}, 0}
	for _, v := range a {
		mh.Push(v)
		fmt.Println(mh.heap)
	}
	res := 0
	for res = mh.Pop(); res != -1; res = mh.Pop() {
		fmt.Println(res)
	}
}
