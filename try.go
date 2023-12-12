package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	result := addTwoNumbers(l1, l2)

	// 打印结果
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}

// 递归方法解决
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwo(l1, l2, 0)
}

// 进位的处理问题 设置carry来进行解决
func addTwo(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	//如果l1、l2都为空，则说明这里两条链表都已经读完
	if l1 == nil && l2 == nil {
		// 最后一位进位处理，若是进位标志为1.创建下一个节点
		if carry > 0 {
			return &ListNode{Val: carry}
		}
		//最后一位不存在进位则直接返回nil
		return nil
	}

	middle := &ListNode{} //实例化middle

	//对不为空节点的进行下一位节点的寻找
	if l1 != nil {
		middle.Val += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		middle.Val += l2.Val
		l2 = l2.Next
	}
	middle.Val += carry

	// 处理进位
	if middle.Val >= 10 {
		middle.Val %= 10
		carry = 1
	} else {
		carry = 0
	}

	// 递归处理下一个节点
	middle.Next = addTwo(l1, l2, carry)
	return middle
}
