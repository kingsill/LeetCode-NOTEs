package main

import "fmt"

func main() {
	L := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	n := 3
	r := removeNthFromEnd(L, n)
	fmt.Println(r)
}

var times int = 0

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

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
}
