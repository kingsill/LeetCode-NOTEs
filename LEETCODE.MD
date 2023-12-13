# 王二的力扣历程
[toc]
## 1. 两数之和 1×
这里我们使用 ``哈希表`` 来解决
>v2, ok := m["pear"]  // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值
```go
package main

import "fmt"

var nums = []int{2, 7, 11, 15}
var target = 9

func main() {
	fmt.Printf("twoSum(nums, target): %v\n", twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	result := map[int]int{} //初始化map的语句，后面的{}的赋初值为空,等同于下一句
	// result := make(map[int]int)

	fmt.Printf("result: %T\n", result)
	
	for i, v := range nums {
		if p, ok := result[target-v]; ok {
			return []int{p, i}
		} 
		result[v] = i
	}
	return nil
}

```

&emsp;&emsp;通过遍历逐步完善map
&emsp;&emsp;创建一个哈希表，对于每一个 x，我们首先查询哈希表中是否存在 target - x，然后将 x 插入到哈希表中，即可保证不会让 x 和自己匹配。
&emsp;&emsp;为了方便搜索，我们将索引值设置为期望target-v，将v的序列保存为后面的值
&emsp;&emsp;使用哈希表，可以将寻找 target - x 的时间复杂度降低到从 O(N) 降低到 O(1)。

## 2.两数相加 1×
&emsp;&emsp;在解决的过程中，遇到的难点在于int的位数，链表的递归查询即建立的问题
 &emsp;&emsp;并未自行解出来

### 链表的相关知识
&emsp;&emsp;链表有时会具有头节点，头节点的指针指向第一个节点的地址，其本身的数据域可以根据自己的选择进行赋值

#### 链表的创建:
``链表的结构定义``一般如下，即由本身的数据和指向下一个节点的指针构成
```go
type ListNode struct {
	Val  int
	Next *ListNode//不能直接赋值listnode，避免产生嵌套引用
}
```
&emsp;&emsp;``链表的创建（每个节点存储一位数字）``，在创建的过程中 我们需要设置中间的可变节点，不然我们可能会丢失对链表的第一个节点的索引，以下的例子中我们就使用middle为中间节点；将head设置为头节点，并代表完整的链表
##### 模拟方式建立
&emsp;&emsp;我们利用``迭代``的方法，只要还存在num，就更新创建一个新的节点
```go
// 尝试建立有头节点的链表，关键在于赋值给middle.next
func CreateList(nums int) *ListNode {
	Head := new(ListNode) //这代表一整个链表，并通过这里的头节点进行标注，方便该链表的引用

	middle := Head //middle视作Head链表的中间节点，其一直改变

	for nums > 0 {
		middle.Next = &ListNode{Val: nums % 10}//头节点赋值方法
		fmt.Printf("middle.Val: %v\n", middle.Val)
		middle = middle.Next
		nums /= 10
	}
	return Head
}
```
##### 链表的**递归创建**
```go
// 递归建立链表
func RecurCreateList(nums int) *ListNode {
	//在递归时好像不需要单独保存头节点位置，后续的位置会递归存储在next中，不用考虑被覆盖的问题
	middle := new(ListNode) //建立头指针，其指针不变

	// 123%10=3
	// 12.3%10=2
	// 1.23%10=1

	// 120%10=0
	// 12 %10=2
	// 1.2%10=1

	if nums > 1 || nums%10 > 0 {
		middle.Val = nums % 10
		fmt.Printf("middle.Val: %v\n", middle.Val)

		if nums > 1 {
			nums /= 10
			middle.Next = RecurCreateList(nums)
		}

	}
	return middle
}
```

#### 链表的读取
##### 遍历读取
&emsp;&emsp;链表读取时我们需要根据是否具有头节点进行一定的调整，下面是使用``遍历（迭代）``进行创建的过程
```go
// 尝试遍历读取链表
func ReadList(L ListNode) {
	middle := L //将头节点赋予这里的中间节点middle

	//循环读取链表的内容
	for middle.Next != nil {
		v := middle.Next.Val //由于我们这里判断的是本身节点是否为空，所以在输出时使用下一节点的值进行输出，避免错过某个值
		fmt.Printf("v: %v\n", v)
		// 	// fmt.Printf("L: %v\n", L)
		middle = *middle.Next
		// 	// fmt.Printf("L: %v\n", L)
	}
}
```
##### 递归读取
下面为使用**递归**进行读取的方法
由于节点的定义过程中使用``内嵌``，在建立相关函数时都`使用指针`比较方便[ 虽然前面都没注意:( ],
```go
// 尝试递归读取链表
func RecurReadList(L *ListNode) {

	fmt.Printf("L.Val: %v\n", L.Val) //打印出此节点中的Val

	//如果本结点的指针不为空，即还有下一个节点，继续读取
	if L.Next != nil {
		RecurReadList(L.Next) //将下个节点的指针传入
	}
	//如果运行到这里，说明指针为空，函数也就到此结束了
}


```

### 真正题目部分
这里选择使用递归方法进行解决
```go
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
//--------------------------------------------------------------------------------------------
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
//--------------------------------------------------------------------------------------
// 尝试递归读取链表
func RecurReadList(L *ListNode) {

	fmt.Printf("L.Val: %v\n", L.Val) //打印出此节点中的Val

	//如果本结点的指针不为空，即还有下一个节点，继续读取
	if L.Next != nil {
		RecurReadList(L.Next) //将下个节点的指针传入
	}
	//如果运行到这里，说明指针为空，函数也就到此结束了
}

```

## 3.无重复字符的最长子串

## 4.找两个正序数组的中位数
在本题中遇到了切片的某些问题，这里进行记录
### 关于切片的扩容问题
&emsp;&emsp;如果在程序中出现下列提示：说明你建立的切片容量不够，超出索引：
>panic: runtime error: index out of range [0] with length 0

&emsp;&emsp;Go 语言中的切片是一种动态数组，它可以自动扩容和缩容以适应不同的数据量。在实际开发中，了解切片的扩容和缩容机制对于优化内存使用和提高程序性能至关重要。
golang中