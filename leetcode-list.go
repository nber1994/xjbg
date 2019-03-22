package main

import "fmt"

//合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//
//示例:
//
//输入:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//输出: 1->1->2->3->4->4->5->6<Paste>
func main() {
	a := &ListNode{1, nil}
	b := &ListNode{4, nil}
	c := &ListNode{5, nil}
	d := &ListNode{1, nil}
	e := &ListNode{3, nil}
	f := &ListNode{4, nil}
	g := &ListNode{2, nil}
	h := &ListNode{6, nil}
	a.Next = b
	b.Next = c
	d.Next = e
	e.Next = f
	g.Next = h
	lists := []*ListNode{a, d, g}
	res := mergeKLists(lists)
	for res != nil {
		//fmt.Println(res.Val)
		res = res.Next
	}

	a = &ListNode{1, nil}
	b = &ListNode{2, nil}
	c = &ListNode{3, nil}
	d = &ListNode{4, nil}
	e = &ListNode{5, nil}
	f = &ListNode{6, nil}
	g = &ListNode{7, nil}
	h = &ListNode{8, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g
	g.Next = h
	res = rotateRight(a, 6)
	for res != nil {
		//fmt.Println(res.Val)
		res = res.Next
	}

	a = &ListNode{1, nil}
	b = &ListNode{1, nil}
	c = &ListNode{1, nil}
	d = &ListNode{2, nil}
	e = &ListNode{2, nil}
	f = &ListNode{3, nil}
	g = &ListNode{4, nil}
	h = &ListNode{4, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g
	g.Next = h
	res = deleteDuplicates1(a)
	for res != nil {
		//fmt.Println(res.Val)
		res = res.Next
	}

	a = &ListNode{1, nil}
	b = &ListNode{2, nil}
	c = &ListNode{3, nil}
	d = &ListNode{4, nil}
	e = &ListNode{5, nil}
	f = &ListNode{6, nil}
	g = &ListNode{7, nil}
	h = &ListNode{8, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g
	g.Next = h
	res = reverseList(a)
	for res != nil {
		//fmt.Println(res.Val)
		res = res.Next
	}

	a = &ListNode{1, nil}
	b = &ListNode{2, nil}
	c = &ListNode{3, nil}
	d = &ListNode{4, nil}
	e = &ListNode{5, nil}
	f = &ListNode{6, nil}
	g = &ListNode{7, nil}
	h = &ListNode{8, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g
	g.Next = h
	res = reverseBetween(a, 1, 1)
	for res != nil {
		//fmt.Println(res.Val)
		res = res.Next
	}

	a = &ListNode{1, nil}
	b = &ListNode{1, nil}
	c = &ListNode{1, nil}
	d = &ListNode{1, nil}
	e = &ListNode{1, nil}
	f = &ListNode{1, nil}
	g = &ListNode{1, nil}
	h = &ListNode{1, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g
	g.Next = h
	res = deleteDuplicates1(a)
	for res != nil {
		//fmt.Println(res.Val)
		res = res.Next
	}

	a = &ListNode{1, nil}
	b = &ListNode{2, nil}
	c = &ListNode{3, nil}
	d = &ListNode{4, nil}
	e = &ListNode{5, nil}
	f = &ListNode{6, nil}
	g = &ListNode{7, nil}
	h = &ListNode{8, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g
	g.Next = h
	h.Next = c
	res = detectCycle(a)
	//fmt.Println(res.Val)

	a = &ListNode{1, nil}
	b = &ListNode{2, nil}
	c = &ListNode{3, nil}
	d = &ListNode{4, nil}
	e = &ListNode{4, nil}
	f = &ListNode{3, nil}
	g = &ListNode{2, nil}
	h = &ListNode{1, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g
	g.Next = h
	o := isPalindrome(a)
	fmt.Println(o)

	linkedList := Constructor()
	linkedList.AddAtHead(5)
	linkedList.AddAtHead(2)
	linkedList.DeleteAtIndex(1)
	linkedList.AddAtIndex(1, 9)
	linkedList.AddAtHead(4)
	linkedList.AddAtHead(9)
	linkedList.AddAtHead(8)
	//fmt.Println(linkedList.Get(3))
	linkedList.AddAtTail(1)
	linkedList.AddAtIndex(3, 6)
	linkedList.AddAtHead(3)

	fmt.Println(">>>>>>>>>>===========================")
	fmt.Println(">>>>>>>>>>===========================")
	fmt.Println(">>>>>>>>>>===========================")
	fmt.Println(">>>>>>>>>>===========================")
	fmt.Println(">>>>>>>>>>===========================")

	a = &ListNode{1, nil}
	b = &ListNode{2, nil}
	c = &ListNode{3, nil}
	d = &ListNode{4, nil}
	e = &ListNode{5, nil}
	f = &ListNode{6, nil}
	g = &ListNode{7, nil}
	h = &ListNode{8, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g
	g.Next = h
	reorderList(a)
	for a != nil {
		//fmt.Println(a.Val)
		a = a.Next
	}

	a = &ListNode{1, nil}
	b = &ListNode{2, nil}
	//c = &ListNode{3, nil}
	//d = &ListNode{4, nil}
	//e = &ListNode{5, nil}
	//f = &ListNode{6, nil}
	//g = &ListNode{7, nil}
	//h = &ListNode{8, nil}
	a.Next = b
	//b.Next = c
	//c.Next = d
	//d.Next = e
	//e.Next = f
	//f.Next = g
	//g.Next = h
	res = reverseKGroup(a, 2)
	for res != nil {
		//fmt.Println(res.Val)
		res = res.Next
	}

	a = &ListNode{1, nil}
	b = &ListNode{2, nil}
	c = &ListNode{3, nil}
	d = &ListNode{4, nil}
	e = &ListNode{5, nil}
	f = &ListNode{6, nil}
	g = &ListNode{7, nil}
	h = &ListNode{8, nil}
	a.Next = c
	c.Next = e
	e.Next = g
	b.Next = d
	d.Next = f
	f.Next = h
	res = mergeList(a, b)
	for res != nil {
		//fmt.Println(res.Val)
		res = res.Next
	}

	a = &ListNode{1, nil}
	b = &ListNode{3, nil}
	c = &ListNode{6, nil}
	d = &ListNode{2, nil}
	e = &ListNode{8, nil}
	f = &ListNode{5, nil}
	g = &ListNode{4, nil}
	h = &ListNode{7, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g
	g.Next = h
	res = sortList(a)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap struct {
	Heap []*ListNode
	Size int
	Lens int //容量
}

func (this *MinHeap) Pop() *ListNode {
	if this.Size == 0 {
		return nil
	}
	res := this.Heap[0]
	this.autoFixDown()
	return res
}

func (this *MinHeap) autoFixDown() {
	this.Heap[0] = this.Heap[this.Size-1]
	this.Heap = this.Heap[:this.Size-1]
	this.Size--
	i := 0
	j := 0
	for i < this.Size {
		child := &ListNode{0, nil}
		hasChild := false
		if i*2+2 < this.Size {
			if this.Heap[i*2+2].Val < this.Heap[i*2+1].Val {
				child = this.Heap[i*2+2]
				j = i*2 + 2
				hasChild = true
			} else {
				child = this.Heap[i*2+1]
				j = i*2 + 1
				hasChild = true
			}
		} else if i*2+1 < this.Size {
			child = this.Heap[i*2+1]
			j = i*2 + 1
			hasChild = true
		}
		if hasChild && this.Heap[i].Val > child.Val {
			this.Heap[i], this.Heap[j] = this.Heap[j], this.Heap[i]
			i = j
		} else {
			break
		}
	}
}

func (this *MinHeap) Push(v *ListNode) bool {
	if this.Size == this.Lens {
		return false
	}
	if v == nil {
		return false
	}
	this.Heap = append(this.Heap, v)
	this.Size++
	this.autoFixUp()
	return true
}

func (this *MinHeap) autoFixUp() {
	i := this.Size - 1
	j := (i - 1) / 2
	for i > 0 {
		if this.Heap[i].Val < this.Heap[j].Val {
			this.Heap[i], this.Heap[j] = this.Heap[j], this.Heap[i]
			i = j
			j = (i - 1) / 2
		} else {
			break
		}
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	newHead := &ListNode{0, nil}
	lens := len(lists)
	mh := &MinHeap{[]*ListNode{}, 0, lens}
	for _, v := range lists {
		mh.Push(v)
	}
	item := mh.Pop()
	newHead.Next = item
	for item != nil {
		if item.Next != nil {
			mh.Push(item.Next)
		}
		item.Next = mh.Pop()
		item = item.Next
	}
	return newHead.Next
}

//旋转链表
//给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
//
//示例 1:
//
//输入: 1->2->3->4->5->NULL, k = 2
//输出: 4->5->1->2->3->NULL
//解释:
//向右旋转 1 步: 5->1->2->3->4->NULL
//向右旋转 2 步: 4->5->1->2->3->NULL
//示例 2:
//
//输入: 0->1->2->NULL, k = 4
//输出: 2->0->1->NULL
//解释:
//向右旋转 1 步: 2->0->1->NULL
//向右旋转 2 步: 1->2->0->NULL
//向右旋转 3 步: 0->1->2->NULL
//向右旋转 4 步: 2->0->1->NULL

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	lens := 1
	end := head
	for end = head; end.Next != nil; end = end.Next {
		lens++
	}
	k = k % lens
	newHead := &ListNode{0, head}
	tail := newHead
	index := lens - k
	for index > 0 {
		tail = tail.Next
		index--
	}
	end.Next = newHead.Next
	newHead.Next = tail.Next
	tail.Next = nil
	return newHead.Next
}

//删除排序链表中的重复元素
//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
//示例 1:
//
//输入: 1->1->2
//输出: 1->2
//示例 2:
//
//输入: 1->1->2->3->3
//输出: 1->2->3

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates1(head *ListNode) *ListNode {
	//if head == nil {
	//	return head
	//}
	//for item := head; item != nil; item = item.Next {
	//	for item.Next != nil && item.Val == item.Next.Val {
	//		item.Next = item.Next.Next
	//	}
	//}
	//return head

	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

//反转链表
//反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	item := head
	next := item.Next
	item.Next = nil
	pre := item
	item = next
	for item != nil {
		next := item.Next
		item.Next = pre
		pre = item
		item = next
	}
	return pre
}

//递归算法
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

//三个指针反转链表
func reverseList(head *ListNode) *ListNode {
	first, ahead, flip := head, head, head
	for ahead.Next != nil {
		flip = ahead.Next
		ahead.Next = flip.Next
		flip.Next = first
		first = flip
	}
	return first
}

//反转链表II
//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
//说明:
//1 ≤ m ≤ n ≤ 链表长度。
//
//示例:
//
//输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{0, head}
	reverHead := newHead
	for i := m - 1; i > 0; i-- {
		reverHead = reverHead.Next
	}
	first, ahead, flip := reverHead.Next, reverHead.Next, reverHead.Next
	time := n - m
	for time > 0 {
		flip = ahead.Next
		ahead.Next = flip.Next
		flip.Next = first
		first = flip
		time--
	}
	reverHead.Next = first
	return newHead.Next
}

//删除排序链表中的重复元素II
//给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
//
//示例 1:
//
//输入: 1->2->3->3->4->4->5
//输出: 1->2->5
//示例 2:
//
//输入: 1->1->1->2->3
//输出: 2->3
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	newHead := &ListNode{}
	newHead.Next = head
	pre, cursor, sentry := newHead, newHead, newHead
	for cursor.Next != nil {
		sentry = cursor.Next
		if sentry.Val != cursor.Val {
			pre = cursor
			cursor = sentry
		} else {
			for sentry.Val == cursor.Val {
				if sentry.Next != nil {
					sentry = sentry.Next
				} else {
					pre.Next = nil
					return newHead.Next
				}
			}
			pre.Next = sentry
			cursor = sentry
		}
	}
	return newHead.Next
}

func deleteDuplicates21(head *ListNode) *ListNode {
	newHead := &ListNode{}
	newHead.Next = head
	pre := newHead
	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur = cur.Next
		}
		if pre.Next == cur {
			pre = pre.Next
		} else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return newHead.Next
}

//回文链表
//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome1(head *ListNode) bool {
	return false
}

//判断列表是否有环
//给定一个链表，判断链表中是否有环。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
//
//
//示例 1：
//
//输入：head = [3,2,0,-4], pos = 1
//输出：true
//解释：链表中有一个环，其尾部连接到第二个节点。
//
//
//示例 2：
//
//输入：head = [1,2], pos = 0
//输出：true
//解释：链表中有一个环，其尾部连接到第一个节点。
//
//
//示例 3：
//
//输入：head = [1], pos = -1
//输出：false
//解释：链表中没有环。
//
//
//
//
//进阶：
//
//你能用 O(1)（即，常量）内存解决此问题吗？
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	walker, runner := head, head
	for runner != nil && runner.Next != nil {
		walker = walker.Next
		runner = runner.Next.Next
		if walker == runner {
			return true
		}
	}
	return false
}

//环形链表II
//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
//说明：不允许修改给定的链表。
//
//
//
//示例 1：
//
//输入：head = [3,2,0,-4], pos = 1
//输出：tail connects to node index 1
//解释：链表中有一个环，其尾部连接到第二个节点。
//
//
//示例 2：
//
//输入：head = [1,2], pos = 0
//输出：tail connects to node index 0
//解释：链表中有一个环，其尾部连接到第一个节点。
//
//
//示例 3：
//
//输入：head = [1], pos = -1
//输出：no cycle
//解释：链表中没有环。
//
//
//
//
//进阶：
//你是否可以不用额外空间解决此题？
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	walker, runner := head, head
	for runner != nil && runner.Next != nil {
		walker = walker.Next
		runner = runner.Next.Next
		if runner == walker {
			walker = head
			for walker != runner {
				walker = walker.Next
				runner = runner.Next
			}
			return walker
		}
	}
	return nil
}

//回文链表
//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//获取中点
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow
	//反转链表
	pre, cursor, sentry := mid, mid, mid
	for cursor.Next != nil {
		sentry = cursor.Next
		cursor.Next = sentry.Next
		sentry.Next = pre
		pre = sentry
	}
	//判断链表是否相同
	item := pre
	for head != mid {
		if head.Val != item.Val {
			return false
		}
		head = head.Next
		item = item.Next
	}
	return true

}

func echoList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

//设计链表
//设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。
//
//在链表类中实现这些功能：
//
//get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
//addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
//addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
//addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。
//deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
//
//
//示例：
//
//MyLinkedList linkedList = new MyLinkedList();
//linkedList.addAtHead(1);
//linkedList.addAtTail(3);
//linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
//linkedList.get(1);            //返回2
//linkedList.deleteAtIndex(1);  //现在链表是1-> 3
//linkedList.get(1);            //返回3
//
//
//提示：
//
//所有值都在 [1, 1000] 之内。
//操作次数将在  [1, 1000] 之内。
//请不要使用内置的 LinkedList 库。
type MyLinkedList struct {
	head    *ListNode
	tail    *ListNode
	size    int
	listMap []*ListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		nil, nil, 0, []*ListNode{},
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.size-1 {
		return -1
	}
	return this.listMap[index].Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{val, this.head}
	this.head = node
	if this.size == 0 {
		this.tail = this.head
	}
	this.listMap = append([]*ListNode{node}, this.listMap...)
	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{val, nil}
	this.listMap = append(this.listMap, node)
	if this.tail == nil {
		this.tail = node
		this.head = node
	} else {
		this.tail.Next = node
		this.tail = node
	}
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
	} else if index == this.size {
		this.AddAtTail(val)
	} else if index > this.size || index < 0 {
		return
	} else {
		node := &ListNode{val, this.listMap[index]}
		this.listMap[index-1].Next = node
		restMap := append([]*ListNode{}, this.listMap[index:]...)
		this.listMap = append(this.listMap[:index], node)
		this.listMap = append(this.listMap, restMap...)
		this.size++
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.size == 0 {
		return
	} else if index < 0 || index > this.size-1 {
		return
	} else {
		if this.size == 1 {
			this.head = nil
			this.tail = nil
			this.listMap = []*ListNode{}
		} else {
			this.listMap[index-1].Next = this.listMap[index].Next
			this.listMap = append(this.listMap[:index], this.listMap[index+1:]...)
		}
		this.size--
	}
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

//重排链表
//给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
//将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//示例 1:
//
//给定链表 1->2->3->4, 重新排列为 1->4->2->3.
//示例 2:
//
//给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	//获取中点
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	restHead := slow.Next
	slow.Next = nil
	//反转第二个链表
	left, cursor, sentry := restHead, restHead, restHead
	for cursor.Next != nil {
		sentry = cursor.Next
		cursor.Next = sentry.Next
		sentry.Next = left
		left = sentry
	}
	for item1, item2 := head, left; item2 != nil; {
		tmp1, tmp2 := item1, item2
		item1, item2 = item1.Next, item2.Next
		tmp2.Next = tmp1.Next
		tmp1.Next = tmp2
	}
}

//k个一组翻转链表
//给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。
//
//k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。
//
//示例 :
//
//给定这个链表：1->2->3->4->5
//
//当 k = 2 时，应当返回: 2->1->4->3->5
//
//当 k = 3 时，应当返回: 3->2->1->4->5
//
//说明 :
//
//你的算法只能使用常数的额外空间。
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 || k == 1 {
		return head
	}
	//获取长度
	lens := 0
	for item := head; item != nil; item = item.Next {
		lens++
	}
	//获取重复的次数
	time := lens / k
	step := k
	newHead := &ListNode{Next: head}
	preCursor, left, cursor, sentry := newHead, head, head, head
	for time > 0 {
		for step > 1 {
			sentry = cursor.Next
			cursor.Next = sentry.Next
			sentry.Next = left
			left = sentry
			step--
		}
		preCursor.Next = left
		if cursor.Next == nil {
			break
		}
		preCursor, left, cursor, sentry = cursor, cursor.Next, cursor.Next, cursor.Next
		time--
		step = k
	}
	return newHead.Next
}

//链表排序
//在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
//
//示例 1:
//
//输入: 4->2->1->3
//输出: 1->2->3->4
//示例 2:
//
//输入: -1->5->3->4->0
//输出: -1->0->3->4->5
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preSlow, fast, slow := head, head, head
	for fast != nil && fast.Next != nil {
		preSlow = slow
		fast = fast.Next.Next
		slow = slow.Next
	}
	fast = preSlow
	fast.Next = nil
	return mergeList(sortList(head), sortList(slow))
}

func mergeList(a, b *ListNode) *ListNode {
	head, tail := b, b
	item1, item2 := a, b.Next
	if a.Val < b.Val {
		head, tail = a, a
		item1, item2 = a.Next, b
	}
	for item1 != nil || item2 != nil {
		if item1 != nil && item2 != nil {
			if item1.Val < item2.Val {
				tail.Next = item1
				tail = item1
				item1 = item1.Next
			} else {
				tail.Next = item2
				tail = item2
				item2 = item2.Next
			}
		} else if item2 == nil {
			tail.Next = item1
			tail = item1
			item1 = item1.Next
		} else if item1 == nil {
			tail.Next = item2
			tail = item2
			item2 = item2.Next
		}
	}
	return head
}

//链表中点
//给定一个带有头结点 head 的非空单链表，返回链表的中间结点。
//
//如果有两个中间结点，则返回第二个中间结点。
//
//
//
//示例 1：
//
//输入：[1,2,3,4,5]
//输出：此列表中的结点 3 (序列化形式：[3,4,5])
//返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
//注意，我们返回了一个 ListNode 类型的对象 ans，这样：
//ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next = NULL.
//示例 2：
//
//输入：[1,2,3,4,5,6]
//输出：此列表中的结点 4 (序列化形式：[4,5,6])
//由于该列表有两个中间结点，值分别为 3 和 4，我们返回第二个结点。
//
//
//提示：
//
//给定链表的结点数介于 1 和 100 之间。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
