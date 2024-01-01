package main

func main() {
	List := Constructor()

	List.AddAtHead(4)
	List.Get(1)
	List.AddAtHead(1)
	List.AddAtHead(5)
	List.DeleteAtIndex(3)
	List.AddAtHead(7)
	List.Get(3)
	List.Get(3)
	List.Get(3)
	List.AddAtHead(1)
	List.DeleteAtIndex(4)

}

// ListNode 定义链表结构
type ListNode struct {
	Val  int
	Next *ListNode
}

/*// 正常建立，不涉及头节点
// MyLinkedList 为了方便操作链表建立的结构体
type MyLinkedList struct {
	size int
	Head *ListNode
}

// Constructor 初始化链表对象 √
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get 获取第index个节点的数据 √
func (this *MyLinkedList) Get(index int) int {
	if index+1 > this.size {
		return -1
	}

	Node := this.Head

	for ; index > 0; index-- {
		Node = Node.Next
	}

	return Node.Val
}

// AddAtHead 将数据为val的节点添加到第一位 √
func (this *MyLinkedList) AddAtHead(val int) {
	Node := &ListNode{Val: val, Next: this.Head}
	this.Head = Node
	this.size++
}

// AddAtTail 将数据为val的节点添加到最后一位 √
func (this *MyLinkedList) AddAtTail(val int) {
	this.size++
	if this.Head == nil {
		this.Head = &ListNode{Val: val}
		return
	}
	Node := this.Head
	for Node.Next != nil {
		Node = Node.Next
	}
	Node.Next = &ListNode{Val: val}
}

// AddAtIndex 将数据为val的节点插入到第index个节点之前；若链表长度为index，则放置到最后一位；若超过链表长度不做修改 √
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.size == index {
		this.AddAtTail(val)
		return
	}
	if this.size < index {
		return
	}
	//有第index节点
	//如果index=0，则需要改动首结点，直接待用addAtHead
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	cur := this.Head
	Node := cur

	//0 1 2 3	;index=2
	for ; index > 1; index-- {
		Node = Node.Next
	}
	Node.Next = &ListNode{Val: val, Next: Node.Next}
	this.Head = cur
	this.size++
}

// DeleteAtIndex 删除第index个节点
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index+1 > this.size {
		return
	}

	// 0 1 2 3	;index=0
	//如果index=0，依旧需要修改头节点
	if index == 0 {
		this.Head = this.Head.Next
		this.size--
		return
	}

	Node := this.Head
	cur := Node

	////0 1 2 3	;index=3,删除最后一个，只需要将最后.next=nil
	//if index+1 == this.size {
	//	for ; index > 1; index-- {
	//		cur = cur.Next
	//	}
	//	cur.Next = nil
	//	this.Head = Node
	//	this.size--
	//	return
	//}
	//
	////0 1 2 3	;index=2
	//for ; index > 1; index-- {
	//	cur = cur.Next
	//}
	//cur.Next = cur.Next.Next
	//this.Head = Node
	//this.size--

	for ; index > 1; index-- {
		cur = cur.Next
	}
	if index+1 == this.size {
		cur.Next = nil
		this.size--
		return
	}
	cur.Next = cur.Next.Next
	this.Head = Node
	this.size--
}*/

//本次目标时建立具有 头节点 的链表(虚拟头节点，但实际就是有头节点)
//链表结构：H* 0 1 2 3
//优势，在对头节点后的节点进行操作会方便

type MyLinkedList struct {
	Size      int
	DummyHead *ListNode
}

// Constructor 使用虚拟头节点时，很关键的一步在在初始化我们的mylinkedlist，我们需要直接对dummyhead分配内存空间，防止后续对首个节点进行操作时出现问题
func Constructor() MyLinkedList {
	return MyLinkedList{DummyHead: &ListNode{}}
}

func (this *MyLinkedList) Get(index int) int {
	//*h 0 1 2 3	;0
	if this.Size <= index {
		return -1
	}

	DummyHead := this.DummyHead
	cur := DummyHead

	for ; index >= 0; index-- {
		cur = cur.Next
	}

	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {

	//if this.Size == 0 {
	//	this.Head = &ListNode{Next: &ListNode{Val: val}}
	//	this.Size++
	//} else {
	//	this.Head.Next = &ListNode{Val: val, Next: this.Head.Next}
	//	this.Size++
	//}

	//newNode := &ListNode{Val: val} // 创建新节点
	//newNode.Next = this.Head.Next  // 新节点指向当前头节点
	//this.Head.Next = newNode       // 新节点变为头节点
	//this.Size++                    // 链表大小增加1

	//上面第一块没想明白，逻辑有问题；第二部分进行整合，本质一样
	this.DummyHead.Next = &ListNode{Val: val, Next: this.DummyHead.Next}
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.Size == 0 {
		this.AddAtHead(val)
		return
	}
	DummyHead := this.DummyHead
	cur := DummyHead
	//h* 0 1 2 3 4	;
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val}
	this.DummyHead = DummyHead
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	//如果index=链表长度，插入到末尾
	if this.Size == index {
		this.AddAtTail(val)
		return
	}
	//如果index>链表长度，什么都不做

	//<长度，插入到index节点前
	//h* 0 1 2 3 4	；2
	if this.Size > index {
		DummyHead := this.DummyHead
		cur := DummyHead

		for ; index > 0; index-- {
			cur = cur.Next
		}
		cur.Next = &ListNode{Val: val, Next: cur.Next}
		this.DummyHead = DummyHead
		this.Size++
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	//h* 0 1 2 3 4	;0

	if this.Size > index {
		DummyHead := this.DummyHead
		cur := DummyHead

		for ; index > 0; index-- {
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
		this.DummyHead = DummyHead
		this.Size--
	}

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
