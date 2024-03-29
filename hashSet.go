package main

import (
	"fmt"
)

type MyHashSet struct {
	hash []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{make([]int, 1000001)}
}

func (this *MyHashSet) Add(key int) {
	this.hash[key] = 1
}

func (this *MyHashSet) Remove(key int) {
	this.hashMap[key] = 0
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.hash[key] == 1
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
