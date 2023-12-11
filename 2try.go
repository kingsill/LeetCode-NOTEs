package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := CreateList(100)
	fmt.Printf("l: %v\n", l)
	// ReadList(l)
}

var Head ListNode //这是一整个链表

// 尝试建立链表
func CreateList(nums int) ListNode {

	middle := new(ListNode) //middle视作Head链表的中间节点，其一直改变
	Head.Next = middle
	for ; nums >= 0; nums /= 10 {
		middle.Val = nums % 10
		middle.Next = middle
	}

	return Head
}

//尝试便利读取链表
func ReadList(L ListNode) {
	middle := L.Next
	for middle.Next != nil {
		v := L.Val
		fmt.Printf("v: %v\n", v)
		// fmt.Printf("L: %v\n", L)
		middle = middle.Next
		// fmt.Printf("L: %v\n", L)
	}
}
