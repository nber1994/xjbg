package main

import (
	"fmt"
)

type Node struct {
	k    int
	v    int
	next *Node
}

type MyHashMap struct {
	ht   []*Node
	mask int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{make([]*Node, 10000), 10000}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	key = this.hash(key)
	node := this.ht[key]
	for ; node != nil; node = node.next {
		if node.k == key {
			node.v = value
			return
		}
	}
	if node == nil {
		this.ht[key] = &Node{key, value, nil}
	} else {
		this.ht[key] = &Node{key, value, this.ht[key]}
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	node := this.ht[this.hash(key)]
	if node == nil {
		return -1
	}
	for ptr := node; ptr != nil; ptr = ptr.next {
		if ptr.k == key {
			return ptr.v
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	key = this.hash(key)
	if this.ht[key] == nil {
		return
	}
	if this.ht[key].k == key {
		this.ht[key] = this.ht[key].next
		return
	}
	for node := this.ht[key]; node.next != nil; node = node.next {
		if node.next.k == key {
			node.next = node.next.next
			return
		}
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % this.mask
}

func main() {
	ht := Constructor()
	ht.Put(1, 1)
	ht.Put(2, 2)
	ht.Put(2, 1)
	fmt.Println(ht.Get(1))
	fmt.Println(ht.Get(2))
	ht.Remove(1)
	fmt.Println(ht.Get(1))
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
