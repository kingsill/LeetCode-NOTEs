package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//暴力解法
/*func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//考虑到从第一个节点就相交的情况，设立虚拟头节点
	DummyHeadA := &ListNode{Next: headA}
	DummyHeadB := &ListNode{Next: headB}

	//使用map进行记录
	var resultA = map[*ListNode]int{}
	var resultB = map[*ListNode]int{}

	//遍历两个链表
	for DummyHeadA.Next != nil || DummyHeadB.Next != nil {

		if DummyHeadA.Next == nil {
			resultB[DummyHeadB.Next] = DummyHeadB.Next.Val
			DummyHeadB = DummyHeadB.Next
		} else if DummyHeadB.Next == nil {
			resultA[DummyHeadA.Next] = DummyHeadA.Next.Val
			DummyHeadA = DummyHeadA.Next
		} else {
			resultA[DummyHeadA.Next] = DummyHeadA.Next.Val
			DummyHeadA = DummyHeadA.Next

			resultB[DummyHeadB.Next] = DummyHeadB.Next.Val
			DummyHeadB = DummyHeadB.Next
		}

		//检查map中是否记录相交节点
		if _, ok := resultA[DummyHeadB]; ok {
			return DummyHeadB
		}
		if _, ok := resultB[DummyHeadA]; ok {
			return DummyHeadA
		}

	}
	return nil
}
*/

// 双指针
/*func getIntersectionNode(headA, headB *ListNode) *ListNode {

	//遍历两链表得到链表长度
	La := 0
	Lb := 0

	curA := headA
	curB := headB

	for curA != nil {
		curA = curA.Next
		La++
	}

	for curB != nil {
		curB = curB.Next
		Lb++
	}

	//将较长的链表偏移到和较短的链表一样长的位置
	if La > Lb {
		for diff := La - Lb; diff > 0; diff-- {
			headA = headA.Next
		}
	}

	if La <= Lb {
		for diff := Lb - La; diff > 0; diff-- {
			headB = headB.Next
		}
	}

	//同级每个节点进行比较，相同则为相交节点
	for headA != nil {

		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next

	}

	return nil
}*/

// 双指针巧解
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
