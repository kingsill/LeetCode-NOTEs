package main

import "fmt"

func main() {
	List := Constructor()

	//fmt.Println(List.Get(1))
	List.AddAtHead(2)
	//this.AddAtIndex(3, 1)
	List.AddAtTail(5)
	//this.DeleteAtIndex(2)
	fmt.Println(List.Head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	size int
	Head *ListNode
}

// Constructor 初始化链表对象
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get 获取第index个节点的数据
func (this *MyLinkedList) Get(index int) int {
	Node := this.Head

	Node = Node.Next
	return 0
}

// AddAtHead 将数据为val的节点添加到第一位 √
func (this *MyLinkedList) AddAtHead(val int) {
	Node := &ListNode{Val: val, Next: this.Head}
	this.Head = Node

}

// AddAtTail 将数据为val的节点添加到最后一位
func (this *MyLinkedList) AddAtTail(val int) {
	Node := this.Head
	if Node == nil {
		Node = &ListNode{Val: val}
		//this.Head = Node
		return
	}
	for Node.Next != nil {
		Node = Node.Next
	}
	Node.Next = &ListNode{Val: val}
}

// AddAtIndex 将数据为val的节点插入到第index个节点之前；若链表长度为index，则放置到最后一位；若超过链表长度不做修改
func (this *MyLinkedList) AddAtIndex(index int, val int) {

}

// DeleteAtIndex 删除第index个节点
func (this *MyLinkedList) DeleteAtIndex(index int) {

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
