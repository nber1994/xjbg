package main

import "fmt"

func main() {
	a := &TreeNode{1, nil, nil}
	b := &TreeNode{2, nil, nil}
	c := &TreeNode{3, nil, nil}
	d := &TreeNode{1, nil, nil}
	e := &TreeNode{2, nil, nil}
	a.Left = b
	a.Right = c
	d.Left = e
	fmt.Println(isSameTree(a, d))
}

//相同的树
//给定两个二叉树，编写一个函数来检验它们是否相同。
//
//如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//
//示例 1:
//
//输入:       1         1
//          / \       / \
//         2   3     2   3
//
//        [1,2,3],   [1,2,3]
//
//输出: true
//示例 2:
//
//输入:      1          1
//          /           \
//         2             2
//
//        [1,2],     [1,null,2]
//
//输出: false
//示例 3:
//
//输入:       1         1
//          / \       / \
//         2   1     1   2
//
//        [1,2,1],   [1,1,2]
//
//		输出: false
//
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	} else {
		leftRes := isSameTree(p.Left, q.Left)
		rightRes := isSameTree(p.Right, q.Right)
		fmt.Println(leftRes && rightRes)
		return leftRes && rightRes
	}
}
