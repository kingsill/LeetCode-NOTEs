package main

import (
	"fmt"
)

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{6, &ListNode{6, nil}}}}}}}
	//head := &ListNode{}
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

//初代版本
/*func removeElements(head *ListNode, val int) *ListNode {
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
}*/

// 新版本 迭代 虚拟头节点
/*func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := new(ListNode) //dummyHead为虚拟头节点
	dummyHead.Next = head      //定义虚拟头节点的下一个节点为原链表第一个节点
	cur := dummyHead

	//虚拟头节点的使用可以不用考虑原链表的头节点的val为Val，不能直接删除的问题
	//我们检查的不是当前节点的val是否为Val，而是当前的next.val，是为了方便删除符合条件的节点；如果选择检查当前节点的val，则不好删除当前节点，大家可以编程体会

	//先定义逻辑
	//首先是循环的定义，只要head.next！=nil，就继续循环

	//如果原链表 循环到的节点的next.val为 给定的Val ，则将head.next->head.next.next,即下一个指针跳过当前节点,且此时这里是不保存下一个节点的，防止下一个节点的val也是Val
	//如果为空 ，则return dummyHead.next

	for cur.Next != nil {

		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}

	return dummyHead.Next

}
*/

// 迭代 直接使用原链表
/*func removeElements(head *ListNode, val int) *ListNode {

	//依旧是先定义逻辑

	//如果是 原链表的头节点为val的话，直接将head=head。next，且当前过程持续，防止头节点后面的节点也为Val
	//这里前置循环 并且要判定head 是否为nil，防止出错

	//这之后再定义cur：=head
	//这里再判定cur.next.val如果等于val，cur.next=cur.next.next（具体原因看上面使用虚拟头节点部分），跳过下一个节点

	//注意：因为要判定不使用虚拟头节点，在循环中要判断循环到的节点是否为nil，防止出错

	for head != nil && head.Val == val {
		head = head.Next
	}
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head

}*/

// 递归
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
