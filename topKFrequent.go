package main

import (
	"fmt"
)

func main() {
	a := []int{5, 4, 4, 4, 4, 1, 1, 3, 3, 3, 6, 7}
	fmt.Println(topKFrequent(a, 3))
}

func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	for _, v := range nums {
		hashMap[v]++
	}
	fmt.Println(hashMap)
	mh := &minHeap{[]int{}, hashMap, 0, k}
	for k, _ := range mh.hashMap {
		mh.Push(k)
	}
	return mh.PopAll()
}

type minHeap struct {
	heap    []int
	hashMap map[int]int
	size    int
	full    int
}

func (this *minHeap) PopAll() []int {
	ret := []int{}
	for res := this.Pop(); res != -1; res = this.Pop() {
		ret = append(ret, res)
	}
	return ret
}

func (this *minHeap) Push(v int) {
	if this.size < this.full {
		this.size++
		this.heap = append(this.heap, v)
		this.fixUp()
		return
	}
	minVal := this.Pop()
	if this.hashMap[minVal] < this.hashMap[v] {
		minVal = v
	}
	this.heap = append(this.heap, minVal)
	this.size++
	this.fixUp()
}

func (this *minHeap) Pop() int {
	if this.size == 0 {
		return -1
	}
	res := this.heap[0]
	this.heap[0] = this.heap[this.size-1]
	this.heap = this.heap[:this.size-1]
	this.size--
	this.fixDown()
	return res
}

func (this *minHeap) fixDown() {
	i := 0
	j := 0
	for i < this.size {
		child := -1
		if 2*i+2 < this.size {
			if this.hashMap[this.heap[2*i+1]] > this.hashMap[this.heap[2*i+2]] {
				j = 2*i + 2
			} else {
				child = this.heap[2*i+1]
				j = 2*i + 1
			}
		} else if 2*i+1 < this.size {
			child = this.heap[2*i+1]
			j = 2*i + 1
		}
		if child > 0 && this.hashMap[this.heap[i]] > this.hashMap[this.heap[j]] {
			this.heap[i], this.heap[j] = this.heap[j], this.heap[i]
			i = j
		} else {
			break
		}
	}
}

func (this *minHeap) fixUp() {
	i := this.size - 1
	j := (i - 1) / 2
	for i > 0 {
		if j > 0 {
			if this.hashMap[this.heap[i]] < this.hashMap[this.heap[j]] {
				this.heap[i], this.heap[j] = this.heap[j], this.heap[i]
				i = j
				j = (i - 1) / 2
			} else {
				break
			}
		} else {
			break
		}
	}
}
