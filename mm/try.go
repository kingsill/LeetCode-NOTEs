package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := CreateList(13)
	fmt.Printf("l: %v\n", l)
	ReadList(*l)
	// fmt.Println(l.Next.Val)
}

// 尝试建立链表
func CreateList(nums int) *ListNode {
	Head := new(ListNode) //这是一整个链表

	middle := Head //middle视作Head链表的中间节点，其一直改变

	for nums > 0 {
		middle.Next = &ListNode{Val: nums % 10}
		fmt.Printf("middle.Val: %v\n", middle.Val)
		middle = middle.Next
		nums /= 10
	}

	return Head
}

// 尝试便利读取链表
func ReadList(L ListNode) {
	middle := *L.Next
	fmt.Printf("middle: %v\n", middle)
	fmt.Println(middle.Val)
	fmt.Println(middle.Next.Val)
	// for middle.Next != nil {
	// 	v := L.Val
	// 	fmt.Printf("v: %v\n", v)
	// 	// 	// fmt.Printf("L: %v\n", L)
	// 	middle = *middle.Next
	// 	// 	// fmt.Printf("L: %v\n", L)
	// }
}
