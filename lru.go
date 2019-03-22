package main

import (
	"fmt"
)

type node struct {
	Key  int
	Val  int
	Pre  *node
	Next *node
}

type LRUCache struct {
	HashMap map[int]*node
	Head    *node
	Tail    *node
	Lens    int
	Full    int
}

func Constructor(capacity int) LRUCache {
	head := &node{0, 0, nil, nil}
	tail := &node{0, 0, nil, nil}
	head.Next = tail
	tail.Pre = head
	return LRUCache{map[int]*node{}, head, tail, 0, capacity}
}

func (this *LRUCache) Put(key int, value int) {
	if _, exist := this.HashMap[key]; exist {
		this.HashMap[key] = value
		return
	}
	addNode := &node{key, value, nil, nil}
	this.HashMap[key] = addNode
	addNode.Next = this.Head.Next
	this.Head.Next.Pre = addNode
	this.Head.Next = addNode
	addNode.Pre = this.Head
	this.Lens++
	//pit(this.Head)
	if this.Lens > this.Full {
		delete(this.HashMap, this.Tail.Pre.Key)
		this.Tail.Pre = this.Tail.Pre.Pre
		this.Tail.Pre.Next = this.Tail
		this.Lens--
	}
}

func pit(lru LRUCache) {
	res := []int{}
	res1 := []int{}
	node := lru.Head
	node1 := lru.Tail
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	for node1 != nil {
		res1 = append(res1, node1.Val)
		node1 = node1.Pre
	}
	fmt.Println(lru.HashMap)
	fmt.Println(res)
	fmt.Println(res1)
}

func (this *LRUCache) Get(key int) int {
	if fNode, exist := this.HashMap[key]; exist {
		fNode.Pre.Next = fNode.Next
		fNode.Next.Pre = fNode.Pre
		fNode.Next = this.Head.Next
		this.Head.Next.Pre = fNode
		fNode.Pre = this.Head
		this.Head.Next = fNode
		return fNode.Val
	}
	return -1
}

func main() {
	obj := Constructor(2)
	obj.Put(2, 1)
	pit(obj)
	obj.Put(2, 2)
	pit(obj)
	fmt.Println(obj.Get(1))
	pit(obj)
	obj.Put(3, 3)
	pit(obj)
	fmt.Println(obj.Get(2))
	pit(obj)
	obj.Put(4, 4)
	pit(obj)
	fmt.Println(obj.Get(1))
	pit(obj)
	fmt.Println(obj.Get(2))
	pit(obj)
	fmt.Println(obj.Get(3))
	pit(obj)
	//fmt.Println(obj.Get(1))
}
