package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 暴力遍历 map哈希表记录
/*func detectCycle(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	table := map[*ListNode]int{}

	//若不成环，会走出循环
	for dummyHead.Next != nil {

		if _, ok := table[dummyHead.Next]; ok {
			return dummyHead.Next
		}
		table[dummyHead.Next] = 1
		dummyHead = dummyHead.Next
	}
	return nil
}
*/

// 双指针法 快慢指针
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	//若有相交节点，则快慢指针一定会相遇
	for fast != nil && slow != nil {
		slow = slow.Next

		//防止fast指针超出链表限制
		if fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next

		//从相遇节点开始寻找相交节点
		if fast == slow {
			//题目要求不修改链表
			p := head
			for p != slow {
				slow = slow.Next
				p = p.Next
			}

			return p

		}
	}
	return nil
}
