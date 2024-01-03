package main

import (
	"fmt"
)

func main() {
	L := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	n := 4
	r := removeNthFromEnd(L, n)
	fmt.Println(r)
}

var times int = 0

type ListNode struct {
	Val  int
	Next *ListNode
}

/*// 本题里面递归的重点在于确定现在是倒数第几个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//为了修改第一个节点的情况，我们需要使用虚拟头节点
	pre := &ListNode{}
	pre.Next = head

	//如果本节点的Next为空，说明递归到链表的最后一个节点，返回，进入计数部分
	if head == nil {
		return head
	}

	//调用自身，修改next节点
	pre.Next = removeNthFromEnd(head.Next, n)
	times++
	if times == n {
		return head.Next
	}

	return head
}*/

// 递归法
/*func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// 创建一个虚拟头节点，简化边界条件处理
	dummy := &ListNode{Next: head}

	//删除倒数第n个节点
	remove(dummy, n)

	return dummy.Next
}

func remove(node *ListNode, n int) int {

	//循环结束条件，找到最后一个节点
	if node == nil {
		return 0
	}

	index := remove(node.Next, n) + 1

	//如果当前节点的下一个节点是目标节点
	if index == n+1 {
		node.Next = node.Next.Next
	}

	return index
}*/

// 双指针 迭代
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	fast, slow := dummyHead, dummyHead

	//fast和slow指针之间的差值
	var diff int

	//循环结束条件，fast移动到最后一个节点
	for fast.Next != nil {
		//如果差值到n，则fast和slow指针都向后偏移即可
		if diff == n {
			fast = fast.Next
			slow = slow.Next
		} else { //差值不到n，则只有fast指针移动，直到满足差值要求
			fast = fast.Next
			diff++
		}

	}

	if diff == n {
		slow.Next = slow.Next.Next
	}

	return dummyHead.Next
}
