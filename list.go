package main

import (
	"fmt"
)

type node struct {
	next  *node
	value int
}

type MyLinkedList struct {
	head *node
	lens int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{&node{nil, 0}, 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	getNode := this.GetNode(index)
	if nil == getNode {
		return -1
	}
	return getNode.value
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := &node{nil, val}
	node.next = this.head.next
	this.head.next = node
	this.lens++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := &node{nil, val}
	tail := this.GetNode(this.lens - 1)
	fmt.Println(this.lens)
	tail.next = node
	this.lens++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.lens {
		this.AddAtTail(val)
		return
	}
	indexNode := this.GetNode(index)
	if nil == indexNode {
		return
	}
	addNode := &node{nil, val}
	preIndexNode := this.GetNode(index - 1)
	preIndexNode.next = addNode
	addNode.next = indexNode
	this.lens++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	delNode := this.GetNode(index)
	if delNode == nil {
		return
	}
	preDelNode := this.GetNode(index - 1)
	preDelNode.next = delNode.next
	this.lens--
}

func (this *MyLinkedList) GetNode(index int) *node {
	if (index < 0 && index != -1) || index > this.lens-1 {
		return nil
	}
	for ptr, i := this.head, -1; ptr != nil; ptr = ptr.next {
		if i == index {
			return ptr
		}
		i++
	}
	return nil
}

func (this *MyLinkedList) pic() {
	for ptr := this.head; ptr != nil; ptr = ptr.next {
		fmt.Println(ptr.value)
	}
}

func main() {
	list := Constructor()
	list.AddAtHead(1)
	list.pic()
	return
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	res := list.Get(1)
	list.DeleteAtIndex(1)
	res = list.Get(1)
	fmt.Println(res)
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
