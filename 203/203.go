package main

import (
	"fmt"
)

func main() {
	//head := &ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}
	head := &ListNode{}
	//head := &ListNode{7, &ListNode{7, &ListNode{7, &ListNode{}}}}
	val := 6

	c := removeElements(head, val)
	fmt.Println(c)
	for c != nil {
		fmt.Println(c.Val)
		c = c.Next
	}

}

// ListNode 链表 节点 定义
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	//遍历节点

	//targetList := new(ListNode) //golang自建类型可以使用make，如slice等；这里我们自定义类型使用new，new返回指针
	targetList := &ListNode{} //直接实例化，可能更快

	middle := targetList //target作为头节点

	for head != nil {
		if head.Val == val { //检测链表中的val是否等于val

		} else { //如果不用删除，直接赋值
			middle.Next = &ListNode{Val: head.Val}
			middle = middle.Next
		}
		head = head.Next
	}

	return targetList.Next //由于这里头节点数据域无用，直接返回第一个节点
}
