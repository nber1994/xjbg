package main

import (
	"fmt"
)

type maxHeap struct {
	heap []int
	size int
}

func (this *maxHeap) Push(v int) {
	this.heap = append(this.heap, v)
	this.size++
	this.fixUp()
}

func (this *maxHeap) Pop() int {
	if 0 == this.size {
		return -1
	}
	res := this.heap[0]
	this.size--
	this.fixDown()
	return res
}

func (this *maxHeap) fixUp() {
	if this.size == 0 {
		return
	}
	i := this.size - 1
	for i > 0 {
		j := (i - 1) / 2
		if j < 0 {
			break
		}
		if this.heap[i] > this.heap[j] {
			this.heap[i], this.heap[j] = this.heap[j], this.heap[i]
			i = j
		} else {
			break
		}
	}
}

func (this *maxHeap) fixDown() {
	if 0 == this.size {
		return
	}
	this.heap[0] = this.heap[this.size]
	this.heap = this.heap[:this.size]
	i := 0
	j := 0
	for i < this.size {
		child := -1
		if i*2+2 < this.size {
			if this.heap[2*i+1] < this.heap[i*2+2] {
				child = this.heap[2*i+2]
				j = 2*i + 2
			} else {
				child = this.heap[2*i+1]
				j = 2*i + 1
			}
		} else if 2*i+1 < this.size {
			child = this.heap[2*i+1]
			j = 2*i + 1
		}
		if child > 0 && this.heap[i] < child {
			this.heap[i], this.heap[j] = this.heap[j], this.heap[i]
			i = j
		} else {
			break
		}
	}
}

func main() {
	a := []int{5, 4, 1, 6, 3, 7}
	mh := &maxHeap{[]int{}, 0}
	for _, v := range a {
		mh.Push(v)
		fmt.Println(mh.heap)
	}
	for res := mh.Pop(); res != -1; res = mh.Pop() {
		fmt.Println(res)
	}
}
