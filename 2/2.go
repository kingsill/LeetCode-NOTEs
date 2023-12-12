/*
	package main

	type ListNode struct {
		Val  int
		Next *ListNode
	}

	func main() {
		var l1, l2 *ListNode
		addTwoNumbers(l1, l2)
	}

	func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
		t1 := RangeList(l1)
		len1 := len(t1)

		t2 := RangeList(l2)
	}

var index int = 0

	func RangeList(l *ListNode) (t []int) {
		if l.Next != nil {
			t[index] = l.Val*10 ^ index
			index++
			l = l.Next
			RangeList(l)
		}
		return t
	}

	func CreateList(t []int, nums int) (result ListNode) {
		var l ListNode
		c := l
		for ; nums > 0; nums-- {
			c.Val = t[nums]
			var new *ListNode
			c.Next = new
		}
		return l
	}
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1, l2 *ListNode
	addTwoNumbers(l1, l2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwo(l1, l2, 0)
}

//问题1 进位的处理问题 设置carry来进行解决
func addTwo(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	// 遍历两个链表的同时进行 结果链表 的建立
	var head ListNode //设置此为结果链表的头节点
	middle := new(ListNode)
	head.Next = middle

	//1.都为nil
	//2.其中之一为nil
	//3.都不为nil

	// 如果l1 和l2的下一个节点不为0，则可以直接加val
	if l1.Next != nil && l2.Next != nil {
		middle.Val = l1.Val + l2.Val + carry
		//进入下一个节点
		l1 = l1.Next
		l2 = l2.Next
	} else {
		if l1.Next == nil && l2.Next == nil { //如果两个next都为nil
			// 如果当前 结果链表 的值大于10，则进行进位，将进位标志carry置为1
			if middle.Val >= 10 {
				middle.Val %= 10
				carry = 1
			}

		} else { // 如果l1，l2中有一个其下节点为nil
			if l1.Next == nil {
				middle.Val = l2.Val + carry
				l2 = l2.Next
			}
			if l2.Next == nil {
				middle.Val = l1.Val + carry
				l1 = l1.Next
			}
		}

	}
	middle.Next = addTwo(l1, l2, carry)
	return &head
}
