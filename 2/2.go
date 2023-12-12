//经典的大失败版本
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
} */

//第二次自我尝试，小丑版本
/*
	package main

	type ListNode struct {
	Val  int
	Next *ListNode
	}

		var l1 = &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	var l2 = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	func main() {
	var l1, l2 *ListNode
	addTwoNumbers(l1, l2)
	}

	//递归方法解决
	func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
		return addTwo(l1, l2, 0)
	}

	//问题1 进位的处理问题 设置carry来进行解决
	func addTwo(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	// 遍历两个链表的同时进行 结果链表 的建立

	middle := new(ListNode)

	//1.都为nil
	//2.其中之一为nil
	//3.都不为nil
	if middle.Val >= 10 {
		middle.Val %= 10
		carry = 1
	} else {
		carry = 0
	}

	// 如果l1 和l2的下一个节点不为0，则可以直接加val
	if l1.Next != nil && l2.Next != nil {
		middle.Val = l1.Val + l2.Val + carry
		//进入下一个节点
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1.Next == nil && l2.Next == nil { //如果两个next都为nil
		// 如果当前 结果链表 的值大于10，则进行进位，将进位标志carry置为1
		if middle.Val >= 10 {
			middle.Val %= 10
			middle.Next = &ListNode{Val: 1}
			return middle
		}
		return nil

	}

	// 如果l1，l2中有一个其下节点为nil
	if l1.Next == nil {
		middle.Val = l2.Val + carry
		l2 = l2.Next
	}
	if l2.Next == nil {
		middle.Val = l1.Val + carry
		l1 = l1.Next
	}

	middle.Next = addTwo(l1, l2, carry)
	return middle
	}

*/

// 大师gpt版本
/* package main

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
*/

// 最后的一次冲刺，冲刺自我版本,堵上尊严的一战！！
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//自定义初始化数据
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	c := addTwoNumbers(l1, l2)
	RecurReadList(c)
}

// 递归方法解决
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwo(l1, l2, 0)
}

func addTwo(l1 *ListNode, l2 *ListNode, carry int) *ListNode {

	middle := &ListNode{}

	//递归传入的指针为空的话，说明计算到最后一位，最后只需要处理进位的问题即可
	if l1 == nil && l2 == nil {
		if carry == 1 {
			middle = &ListNode{Val: 1}
			return middle
		}
		return nil
	}

	//处理l1、l2任意一个不为1的情况
	if l1 != nil {
		middle.Val += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		middle.Val += l2.Val
		l2 = l2.Next
	}
	middle.Val += carry
	//处理计算到最后的进位处理问题
	if middle.Val >= 10 {
		carry = 1
		middle.Val %= 10
	} else {
		carry = 0
	}

	middle.Next = addTwo(l1, l2, carry)
	return middle

}

// 尝试递归读取链表
func RecurReadList(L *ListNode) {

	fmt.Printf("L.Val: %v\n", L.Val) //打印出此节点中的Val

	//如果本结点的指针不为空，即还有下一个节点，继续读取
	if L.Next != nil {
		RecurReadList(L.Next) //将下个节点的指针传入
	}
	//如果运行到这里，说明指针为空，函数也就到此结束了
}

//无敌的收尾！！！！！！！！！！！！！
