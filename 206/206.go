package main

import (
	"fmt"
)

func main() {
	origin := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	t1 := &ListNode{}
	result := reverseList(origin)
	result2 := reverseList(t1)
	fmt.Println("ok", result, result2)
	pre := &ListNode{}
	fmt.Println(pre)
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 初始方法，建立新链表
/*func reverseList(head *ListNode) *ListNode {
	NewNode := &ListNode{}

	for ; head != nil; head = head.Next {
		NewNode.Val = head.Val
		NewNode = &ListNode{Next: NewNode}
	}

	return NewNode.Next
}
*/

// 迭代 双指针法

//func reverseList(head *ListNode) *ListNode {
//
//	////如果原链表为空或者nil，直接返回原链表
//	//if head == nil || head.Next == nil {
//	//	return head
//	//}
//
//	/*	// 0 1 2 3 4
//		cur := head
//
//		//pre := &ListNode{}
//		var pre *ListNode //这样定义不分配内存，在后续第一次向外赋值即为nil；而上面的定义则会为0
//
//		//结果后续函数可以覆盖前面的特殊情况 ：）
//		for cur != nil {
//			c := cur.Next
//			cur.Next = pre
//			pre = cur
//			cur = c
//		}
//		return pre*/
//
//	//事实上可以直接操作原链表，不用建立副本
//
//	var pre *ListNode //这样定义不分配内存，在后续第一次向外赋值即为nil；而上面的定义则会为0
//
//	//结果后续函数可以覆盖前面的特殊情况 ：）
//	for head != nil {
//		c := head.Next
//		head.Next = pre
//		pre = head
//		head = c
//	}
//	return pre
//}

// 递归，也为双指针
// 把自己和后面的换
/*func reverseList(head *ListNode) *ListNode {
	//递归结束条件
	if head == nil || head.Next == nil {
		return head
	}

	//调用自身
	newHead := reverseList(head.Next)

	//处理递归
	//由于翻转链表，那么我们的目的就是目前节点的下一个改为上一个
	//目前节点head.next
	head.Next.Next = head
	//上面这么写会使链表成环，这里要打破 环
	head.Next = nil

	return newHead
}
*/

// 重来
func reverseList(head *ListNode) *ListNode {
	return reverse(head)
}

func reverse(head *ListNode) *ListNode {
	Node := head

	//结束循环的条件 节点到达尾节点
	if Node.Next.Next != nil {
		reverse(Node.Next)
	}

	Node.Next.Next = Node
	return Node.Next
}
