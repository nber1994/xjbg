package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack struct {
	BinTree []*TreeNode
	Top     int
}

func (nodeStack *stack) push(node *TreeNode) {
	nodeStack.BinTree = append(nodeStack.BinTree, node)
	nodeStack.Top++
}

func (nodeStack *stack) pop() *TreeNode {
	var res *TreeNode
	nodeStack.Top--
	if nodeStack.Top < 0 {
		res = nil
	} else {
		res = nodeStack.BinTree[nodeStack.Top]
	}
	nodeStack.BinTree = nodeStack.BinTree[:nodeStack.Top]
	return res
}

func inorderTraversal1(root *TreeNode) []int {
	res := []int{}
	myStack := &stack{[]*TreeNode{}, 0}
	for nil != root || myStack.Top != 0 {
		for nil != root {
			myStack.push(root)
			root = root.Left
		}
		root = myStack.pop()
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	myStack := &stack{[]*TreeNode{}, 0}
	for nil != root || myStack.Top != 0 {
		for nil != root {
			myStack.push(root)
			root = root.Left
		}
		root = myStack.pop()
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	lens = len(inorder)
	mid := 0
	for i := 0; i < lens; i++ {
		if inorder[i] == postorder[lens-1] {
			mid = i
		}
	}
	return &TreeNode{inorder[mid].Val, buildTree(inorder[:mid], postorder[:mid]), buildTree(inorder[mid+1:lens], postorder[mid:lens-1])}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}
	return isSym(root.Left, root.Right)
}

func isSym(left, right *TreeNode) bool {
	if nil == left && nil == right {
		return true
	}
	if nil == left || nil == right || left.Val != right.Val {
		return false
	}
	outerRes := isSym(left.Left, right.Right)
	innerRes := isSym(left.right, right.Left)
	return outerRes && innerRes
}

func isSymmetric1(root *TreeNode) bool {
	s := *stack{0, nil, nil}
	if root == nil {
		return true
	}
	stack.push(root.Left)
	stack.push(root.Right)
	for s.Top > 0 {
		left := s.pop()
		right := s.pop()
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val == right.Val {
			return false
		}
		s.push(left.Left)
		s.push(right.Right)
		s.push(left.Right)
		s.push(right.Left)
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	leftD := maxDepth(root.Left)
	rightD := maxDepth(root.Right)
	if leftD > rightD {
		return leftD + 1
	} else {
		return rightD + 1
	}
}

func maxDepth1(root *TreeNode) int {
	deps := 0
	que := []*TreeNode{root}
	lens := len(que)
	for lens > 0 {
		for i := 0; i < lens; i++ {
			if nil == que[i] {
				continue
			}
			que = append(que, que[i].Left)
			que = append(que, que[i].Right)
		}
		que = que[lens:]
		lens = lens(que)
		if lens > 0 {
			deps++
		}
	}
	return deps
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if nil == root {
		return res
	}
	que := []*TreeNode{root}
	lens := len(que)
	for lens > 0 {
		level := []int{}
		for i := 0; i < lens; i++ {
			level = append(level, que[i].Val)
			if nil != que[i].Left {
				que = append(que, que[i].Left)
			}
			if nil != que[i].Right {
				que = append(que, que[i].Right)
			}
		}
		que = que[lens:]
		lens = len(que)
		res = append([][]int{level}, res)
	}
	return res
}

func levelOrderBottom1(root *TreeNode) [][]int {
	res := [][]int{}
	res = levelOrder(0, root, res)
}

func levelOrder(level int, root *TreeNode, res [][]int) [][]int {
	if nil == root {
		return res
	}
	if _, exist := res[level]; !exist {
		res = append(res, []int)
	}
	res[level] = append(res[level], root.Val)
	res = levelOrder(level+1, root.Left, res)
	return levelOrder(level+1, root.Right, res)
}

func minDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 {
		return right + 1
	} else if right == 0 {
		return left + 1
	} else {
		if left > right {
			return right + 1
		} else {
			return left + 1
		}
	}
}

func isBalanced(root *TreeNode) bool {
	res := false
	dep := isBlan(root)
	if dep > 0 {
		res = true
	}
	return res
}

func isBlan(root *TreeNode) int {
	if nil == root {
		return 0
	}
	left := isBlan(root.Left)
	right := isBlan(root.Right)
	if left < 0 || right < 0 {
		return -1
	}
	diff := left - right
	if diff < -1 || diff > 1 {
		return -1
	} else if diff >= 0 {
		return left + 1
	} else {
		return right + 1
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	return recursive(0, root, res)
}

func recursive(dep int, root *TreeNode, res [][]int) [][]int {
	if nil == root {
		return res
	}
	if dep+1 > lens(res) {
		res = append(res, []int)
	}
	if dep/2 == 0 {
		res[dep] = append(res[dep], root.Val)
	} else {
		res[dep] = append([]int{root.Val}, res[dep]...)
	}
	res := recursive(dep+1, root.Left, res)
	return recursive(dep+1, root.Right, res)
}

func zigzagLevelOrder1(root *TreeNode) [][]int {
	res := [][]int{}
	que := []*TreeNode{root}
	lens := len(que)
	dep := 0
	for lens > 0 {
		level := []int{}
		for i := 0; i < lens; i++ {
			if nil == que[i] {
				continue
			}
			if dep%2 == 0 {
				level = append(level, que[i].Val)
			} else {
				level = append([]int{que[i].Val}, level...)
			}
			que = append(que, que[i].Left)
			que = append(que, que[i].Right)
		}
		res = append(res, level)
		que = que[lens:]
		lens = len(que)
		dep++
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	head := preorder[0]
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == head {
			index = i
		}
	}
	return &TreeNode{root.Val, buildTree(preorder[1:index+1], inorder[:index]), buildTree(preorder[index+1:], inorder[index+1:])}
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	sum -= sum - root.Val
	if root.Left != nil {
		leftRes := hasPathSum(root.Left, sum)
	}
	if root.Left != nil {
		rightRes := hasPathSum(root.Right, sum)
	}
	return leftRes && rightRes
}

func insertionSortList(head *ListNode) *ListNode {
	myHead := &ListNode{0, nil}
	for head != nil {
		node := head
		for nowHead := myHead; nowHead.Next != nil; nowHead = nowHead.Next {
			if nowHead.Next.Val > node.Val {
				node.Next = nowHead.Next
				nowHead.Next = node
				break
			}
		}
	}
	return myHead
}

func rightSideView(root *TreeNode) []int {
	que := []*TreeNode{root}
	lens := 1
	res := []int{}
	for lens > 0 {
		level := []int{}
		for i := 0; i < lens; i++ {
			if que[i] == nil {
				continue
			}
			level = append(level, que[i].Val)
			que = append(que, que[i].Left)
			que = append(que, que[i].Right)
		}
		res = append(res, level[lens-1])
		que = que[:lens]
		lens = len(que)
	}
	return res
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	if nil == root {
		return 0
	}
	if root.Val >= L {
		sum += rangeSumBST(root.Left, L, R)
	}
}
