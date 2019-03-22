package main

import (
	"fmt"
)

func main() {
	a := &TreeNode{7, nil, nil}
	b := &TreeNode{3, nil, nil}
	c := &TreeNode{15, nil, nil}
	d := &TreeNode{9, nil, nil}
	e := &TreeNode{20, nil, nil}
	a.Left = b
	a.Right = c
	c.Left = d
	c.Right = e
	obj := Constructor(a)
	for obj.HasNext() {
		obj.Next()
		//fmt.Println(obj.Next())
	}
	fmt.Println(scoreOfParentheses("(()(()))"))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
//
//调用 next() 将返回二叉搜索树中的下一个最小的数。
//
//BSTIterator iterator = new BSTIterator(root);
//iterator.next();    // 返回 3
//iterator.next();    // 返回 7
//iterator.hasNext(); // 返回 true
//iterator.next();    // 返回 9
//iterator.hasNext(); // 返回 true
//iterator.next();    // 返回 15
//iterator.hasNext(); // 返回 true
//iterator.next();    // 返回 20
//iterator.hasNext(); // 返回 false
//
//
//提示：
//
//next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
//你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode
	size  int
}

func (this *BSTIterator) Push(v *TreeNode) {
	this.stack = append(this.stack, v)
	this.size++
}

func (this *BSTIterator) Pop() *TreeNode {
	if this.size == 0 {
		return nil
	}
	res := this.stack[this.size-1]
	this.stack = this.stack[:this.size-1]
	this.size--
	return res
}

func Constructor(root *TreeNode) BSTIterator {
	myBSTIterator := new(BSTIterator)
	myBSTIterator.stack = []*TreeNode{}
	myBSTIterator.size = 0
	for root != nil {
		myBSTIterator.Push(root)
		root = root.Left
	}
	return *myBSTIterator
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	res := this.Pop()
	item := res.Right
	for item != nil {
		this.Push(item)
		item = item.Left
	}
	return res.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.size > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

//括号的分数
//给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：
//
//() 得 1 分。
//AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
//(A) 得 2 * A 分，其中 A 是平衡括号字符串。
//
//
//示例 1：
//
//输入： "()"
//输出： 1
//示例 2：
//
//输入： "(())"
//输出： 2
//示例 3：
//
//输入： "()()"
//输出： 2
//示例 4：
//
//输入： "(()(()))"
//输出： 6
//
//
//提示：
//
//S 是平衡括号字符串，且只含有 ( 和 ) 。
//2 <= S.length <= 50
type Stack struct {
	stack []int
	size  int
}

func (this *Stack) push(v int) {
	this.stack = append(this.stack, v)
	this.size++
}

func (this *Stack) pop() int {
	res := this.stack[this.size-1]
	this.stack = this.stack[:this.size-1]
	this.size--
	return res
}

func scoreOfParentheses(S string) int {
	st := &Stack{[]int{}, 0}
	sa := []byte(S)
	for _, v := range sa {
		if v == '(' {
			st.push('(' - '0')
		} else {
			pre := st.pop()
			if pre == '('-'0' {
				st.push(1)
			} else {
				tmp := []int{}
				for next := pre; next != '('-'0'; next = st.pop() {
					tmp = append(tmp, next)
				}
				st.push(2 * arraySum(tmp))
			}
		}
	}
	return arraySum(st.stack)
}

func arraySum(nums []int) int {
	sum := 0
	for _, v := range nums {
		if v > 0 {
			sum += v
		}
	}
	return sum
}
