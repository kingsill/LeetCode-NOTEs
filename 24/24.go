package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	L := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}

	h := swapPairs(L)
	fmt.Printf("h: %v\n", h)
}

//递归 每两个进行一次交换
/*func swapPairs(head *ListNode) *ListNode {
	//每两个的第一个节点
	var pre *ListNode
	//递归结束条件，即当剩下的节点不满足两个时
	if head == nil || head.Next == nil {
		return head
	}

	//保留递归的子节点
	cur := head.Next.Next

	//交换本次的顺序
	pre = head.Next
	pre.Next = head

	//1 2 3 4 -> 2 1 4 3
	//第三个节点开始进行递归，且如果交换，指向的应该是交换过后的次序，即本来的第四个
	pre.Next.Next = swapPairs(cur)

	//将交换次序后的返回
	return pre

}*/

// 使用迭代方法
// 不能操之过急，这里前面迭代完连接的应当还是未交换过的次序
func swapPairs(head *ListNode) *ListNode {

	dummyHead := &ListNode{}
	cur := dummyHead //用来保存当前位置

	cur.Next = head
	for cur.Next != nil && cur.Next.Next != nil {

		node1 := cur.Next
		node2 := cur.Next.Next
		next := cur.Next.Next.Next

		cur.Next = node2
		node2.Next = node1
		node1.Next = next
		cur = node1
	}

	return dummyHead.Next
}
