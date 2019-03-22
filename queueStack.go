package main

import (
	"fmt"
)

type MyStack struct {
	queue1 []int
	queue2 []int
	lens   int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{[]int{}, []int{}, 0}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue2 = append(this.queue2, this.queue1...)
	fmt.Println(this.queue1)
	this.queue1 = this.queue1[0:0]
	this.queue1 = append(this.queue1, x)
	this.queue1 = append(this.queue1, this.queue2...)
	this.queue2 = this.queue1[0:0]
	fmt.Println("-===========-")
	this.lens++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	res := this.queue1[0]
	this.queue1 = this.queue1[1:]
	this.lens--
	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	res := this.queue1[0]
	return res
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.lens == 0
}

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
