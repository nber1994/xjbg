package main

import (
	"fmt"
)

type MinStack struct {
	stack []int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, 0}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.min = x
	}
	this.stack = append(this.stack, x-this.min)
	if x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	tail := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if tail < 0 {
		this.min = this.min - tail
	}
}

func (this *MinStack) Top() int {
	tail := this.stack[len(this.stack)-1]
	if tail >= 0 {
		return tail + this.min
	} else {
		return this.min
	}
}

func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	minS := Constructor()
	minS.Push(1)
	minS.Push(2)
	fmt.Println(minS.stack)
	res := minS.Top()
	fmt.Println(res)
	res = minS.GetMin()
	fmt.Println(res)
	minS.Pop()
	fmt.Println(minS.stack)
	minS.GetMin()
	fmt.Println(res)
	res = minS.Top()
	fmt.Println(res)
}

/**
* Your MinStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.GetMin();
 */
