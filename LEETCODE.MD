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

## 3.无重复字符的最长子串 1×√
### 我的理解 第一版
```go
func lengthOfLongestSubstring(s string) int {
	//设置为每次保存最大个数的窗口
	//type Win map[string]int

	//每次循环的窗口
	n := make(map[string]int)

	//保存每次循环得到的最大值
	var Max int

	//记录每个窗口中最多的字母数
	i := 0
	var light int = 1     //用作窗口的左侧指针
	for c, v := range s { //c+1用来记录找到第几个位置
		//fmt.Println(string(v))
		fmt.Println(n)
		fmt.Println("c+1(n):", c+1)

		_, ok := n[string(v)]

		//fmt.Println(ok)

		if ok { //如果有对应的字母，说明有重复
			//先进行计算，再进行下一个窗口的布置
			if n[string(v)] >= light { //如果这个重复的字母在light的右边，则需要用这个字母的位置进行最长长度的jisuan
				fmt.Println("一")
				i = c + 1 - n[string(v)]
			} else { //否则用light计算即可
				fmt.Println("二")
				i = c + 1 - light
			}

			light = max(light, n[string(v)]) //更新light的位置
			n[string(v)] = c + 1             //更新重复字母最后的位置
			fmt.Println("light:", light)

			fmt.Println("i:", i)

		} else { //如果没有对应的字母1.最长长度计数+1 2.将临时最长长度长度保存

			n[string(v)] = c + 1
			i++
		}

		//在遍历的过程中不断更新最长长度，防止最后一段字母是最长的
		if i >= Max {
			Max = i
			fmt.Println("MAX:", Max)
		}
	}
	return Max
}
```

### 进阶学习版本，即官方版本
```go
```
#### 滑动窗口以及map的相关知识
&emsp;&emsp;**滑动窗口(同向双指针)** 常用于 **子串** 以及 **子数组** 等相关格式
&emsp;&emsp;利用for range可以遍历字符串、数组等，得到的index可以作为**右指针**；**左指针**则根据要求进行调整即可‘
>map中delete函数的使用
	**func delete(m map[Type]Type1, key Type)**
删除m中字段为key的一条

### 自我反思 
&emsp;&emsp;第一版解法相对于官方的问题所在：
1.对于左指针的理解不够，没有理解滑动窗口的*滑动*精髓
2.没有利用好range过程中产生的排序

## 4.找两个正序数组的中位数 1√
在本题中遇到了切片的某些问题，这里进行记录
### 关于切片的扩容问题
&emsp;&emsp;如果在程序中出现下列提示：说明你建立的切片容量不够，超出索引：
>panic: runtime error: index out of range [0] with length 0

####  切片的创建
&emsp;&emsp;Go 语言中的切片是一种动态数组，它可以自动扩容和缩容以适应不同的数据量。在实际开发中，了解切片的扩容和缩容机制对于优化内存使用和提高程序性能至关重要。
切片的创建方式一般来说有三种
```go
//1.通过数组的切割进行切片的建立
nums:=[5]int{1,2,3}//初始化一个切片实例
slice：=nums[:]//注意切割无论是切割数组还是切片都是**``左闭右开``**
//var slice []int

//2.声明切片并初始化
//2.1
var slice []int//这种方式初始化的切片并未自动分配内存，**``长度和容量都为0``**
//2.2
slice:=[]int{1,2,3}//声明并初始化

//3.通过make函数进行建立 推荐※※
slice:=make([]int,3,5)//第一个参数是长度，第二个参数是容量

```

#### 切片的扩容
&emsp;&emsp;golang中当向切片中添加元素时，如果切片的容量不足，就会触发切片的扩容机制。

>切片的扩容会分配一个新的数组，并将原来的元素复制到新数组中。

&emsp;&emsp;新数组的长度通常是原数组长度的两倍，但是如果原数组长度小于 1024，新数组的长度会比原数组长度大 1 倍；如果原数组长度大于等于 1024，新数组的长度会比原数组长度大 1/4 倍。
&emsp;&emsp;切片的扩容可以通过内置函数 append 来实现。当切片容量不足时，调用 append 函数会触发切片的扩容。

>append函数签名如下：
**func Append(slice []int, items ...int) []int**

#### 切片的缩容
&emsp;&emsp;当从切片中删除元素时，如果切片的长度小于容量的 1/4，就会触发切片的缩容机制。切片的缩容会释放底层数组中未使用的空间。新数组的长度通常是原数组长度的一半，但是如果原数组长度小于 1024，新数组的长度等于原数组长度；如果原数组长度大于等于 1024，新数组的长度会比原数组长度小 1/2 倍。

&emsp;&emsp;切片的缩容不需要显式调用任何函数，Go 语言的垃圾回收机制会自动进行。
&emsp;&emsp;切片的缩容使用我们前面提到的**切割**即可实现

### 二分查找法(第k小数)<span id="jump2"><span>
[更加详细部分](#jump)
&emsp;&emsp;二分查找用于在**有序**且**不重复**的元素列表中寻找需要的元素，返回其位置或错误

&emsp;&emsp;当要求算法的时间复杂度在**O（logn）**等带log的复杂度时，可以考虑二分查找法

由于在go中int类型计算时如果带小数，则取整，如``k=5;k/2=2``。通过此我们可以通过for循环来得到第k小数
```go
//如果我们这里寻找第9小的数，前面有8个位置
k:=9
for{
	k=k-k/2
	if k==1{
		break
	}
} 
// 9-9/2=5 4
// 5-5/2=3 2 
// 3-3/2=2 1
// 2-2/1=1 1
```


### 二分查找法我的代码
```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	z := len(nums1) + len(nums2) //两个数组的总的长度
	if z%2 == 1 {                //奇数
		return float64(Get(nums1, nums2, (z+1)/2)) //奇数的中位数为中间的一个数
	} else { //偶数
		return (float64(Get(nums1, nums2, z/2) + Get(nums1, nums2, z/2+1))) / 2 //偶数的中位数为中间二数的平均值
	}
}

// Get 作用为获取nums1和nums2中的第k小数
func Get(nums1 []int, nums2 []int, k int) int {
	var Rindex1, Rindex2 = 0, 0 //为我们总的偏移量
	offset1, offset2 := 0, 0    //当前一次循环的偏移量

	var flow int = 0 //定义flow为我们要求的数组的查询位置超限

	for {
		//如果nums1或2已经查询完毕或者长度为0
		if Rindex1 == len(nums1) {
			return nums2[k+Rindex2-1]
		}
		if Rindex2 == len(nums2) {
			return nums1[k+Rindex1-1]
		}

		//k==1说明已经查询完毕，返回后一个数即为第k小数
		if k == 1 {

			return min(nums1[Rindex1], nums2[Rindex2])
		}

		//计算下一次循环的指针位置（没有越限情况）
		point1 := k/2 - 1 + Rindex1
		point2 := k/2 - 1 + Rindex2

		//判断有没有溢出，进行越位处理
		if k/2-1+Rindex1 > len(nums1)-1 {
			point1 = len(nums1) - 1
			flow = 1
		}
		if k/2-1+Rindex2 > len(nums2)-1 {
			point2 = len(nums2) - 1
			flow = 2
		}

		if flow == 1 || flow == 2 {
			if flow == 1 { //如果时数组1超限
				if nums1[point1] <= nums2[point2] {
					offset1 = len(nums1)
					k = k - len(nums1)
				} else {
					offset2 = k / 2
					k -= k / 2
				}

			}
			if flow == 2 { //如果是数组2超限
				if nums1[point1] <= nums2[point2] {
					offset1 = k / 2
					k -= k / 2
				} else {
					offset2 = len(nums2)
					k = k - len(nums2)
				}

			}

		} else {
			if nums1[point1] <= nums2[point2] {
				offset1 = k / 2
			} else {
				offset2 = k / 2
			}

			k -= k / 2
		}
		//计算总偏移量
		Rindex1 += offset1
		Rindex2 += offset2
		//重置单次偏移量及标志位
		offset1, offset2, flow = 0, 0, 0
	}

}
```
#### 官方版本
&emsp;&emsp;与官方相比，对与越限的处理不够简单
```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
```



## 5.最长回文子串 1×
### 我的思路
1. 从两侧开始找相同的字符，再向内寻找相同的字符，但是在如何循环进行查找无法实现
2. 每个字符查看有无相近的两个，从每个这样的实例开始向外扩散求解，也是在循环的建立以及每个字符相近部分的指针保存上有问题
&emsp;&emsp;简而言之，上面两种思路我都没有实现

### 正确解法
1. 动态规划
1.1 什么是动态规划
&emsp;&emsp;通过组合**子问题**的解来求解原问题，与**分治**方法类似。分治方法将问题划分为互不相交的子问题，递归求解子问题，再组合。
&emsp;&emsp;而我们这里的**动态规划**是子问题具有公共的子问题。通过其子子问题的解影响子问题。
1.2 动态规划的应用场景
&emsp;&emsp;动态规划方法一般用来求解**最优化问题**。这类问题可以有很多可行解，每个解都有一个值，我们希望找到具有最优值的解，我们称这样的解为问题的一个最优解，而不是最优解，因为可能有**多个解都达到最优值**。
1.3 动态规划的步骤
&emsp;&emsp;1、定义一个状态，这是一个最优解的结构特征
&emsp;&emsp;2、进行状态递推，得到递推公式---我们这里的 **动态转移方程** 即为 子子串为回文 **且** 本次两个字符相等
&emsp;&emsp;3、进行初始化
&emsp;&emsp;4、返回结果
&emsp;&emsp;每次计算将其解保存在表格中，后续求解进行查询
1.4 具体实现
&emsp;&emsp;由于回文子串在去掉最外面的两个字符之后依旧是回文子串，所以非常典型地满足回文子串的特性。
&emsp;&emsp;用于参考的图片：
![Alt text](leetcode5动态规划.png)


**思路：**
&emsp;&emsp;依旧是从头开始遍历字符串的每个可能为中心的字符，利用动态规划寻找改字符为中心的最长回文。比较后返回整体最长的回文子串。

**代码实现：**
&emsp;&emsp;
```go
func main() {
	s := "aacabdkacaa"
	r := longestPalindrome(s)
	fmt.Println(r)
}

func longestPalindrome(s string) string {
	length := len(s) //字符串长度
	var X, Y, M = 0, 0, 0
	table := make([][]int, length) //用于保存信息的表格
	for i := 0; i < length; i++ {
		table[i] = make([]int, length)
	}
	//fmt.Println(table)

	////首先将对角线的元素都标记为回文字符串
	//for i := 0; i < length; i++ {
	//	table[i][i] = 0
	//}
	//不需要以上定义，初始化表格的过程中自动赋值为0
	//为了方便，我们直接通过初始化将所有标记为0，意味都为回文子串，通过挨个修正来得到正确结果
	//其目的是为了省略对没有左下角数据的字符的额外判断

	//从第一行开始遍历
	for x := 0; x < length; x++ {
		for y := x + 1; y < length; y++ {
			u := y - x - 1 //大X为实际对应的表格内坐标
			//状态转移方程中的对于子子串的判断在没有左下角的数据时是不启用的
			if s[u] == s[y] && table[u+1][y-1] == 0 { //如果是回文子串，记录其长度
				if y-u >= M { //判断当前回文子串是否为最长回文子串
					M = y - u
					X = u
					Y = y
				}
				//table[u][y] = 0
			} else {
				table[u][y] = 1
			}

		}
	}
	//for _, ints := range table {
	//	fmt.Println(ints)
	//}
	n := s[X : Y+1]
	return n
}

```

## 代码随想录大时代
### 数组
#### 数组理论基础
数组的关键：
>数组下标从0开始
数组内存空间地址是连续的

#### 二分查找 leetcode704<span id="jump"><span>
[前面部分第4题，包括使用条件等](#jump2)
[感谢代码随想录：）](https://www.programmercarl.com/0704.%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE)
&emsp;&emsp;二分查找法中对于 **区间** 的定义

&emsp;&emsp;二分查找涉及的很多的**边界条件**，逻辑比较简单，但就是写不好。例如到底是 ``while(left < right)`` 还是 ``while(left <= right)``，到底是``right = middle``呢，还是要``right = middle - 1``呢？

&emsp;&emsp;对于二分法的区间定义，可以有闭区间、左闭右开、左开右闭等多种写法，他们的含义是我们寻找的**target**与左右指针的关系
与此相关的自然就会产生**循环条件的判断** 以及 **每次循环的左右指针的偏移** 的定义问题
以下为闭区间版本的写法：
```go
// 学习版本二分查找
func search(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	for low <= high { //在闭区间中如[1，1]是有意义的
		mid := low + (high-low)/2 //这样定义防止数字太大溢出
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target { //缩小target所在的区间，由于为闭区间，可以舍弃上次的边界值
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

```
左闭右开区间版本：
```go
// 左闭右开区间版本
func search(nums []int, target int) int {
	high := len(nums) //由于右开，取不到最右边的数，所以这里不用-1也不会超限
	low := 0
	for low < high { //类似于[1,1)这样的区间是没有意义的
		mid := low + (high-low)/2 //同样防止溢出
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid //右开，不需要舍弃右边值
		} else {
			low = mid + 1 //左开，当前值不符合，抛弃当前值
		}
	}
	return -1
}

```

#### 移除元素 leetcode27
&emsp;&emsp;go中数据类型的分类：
1.**值**类型：int、float、bool、string、数组、结构体
2.**引用**类型：指针、切片、map、管道、接口

&emsp;&emsp;由于切片为引用类型，其本质是基于数组完成的，因此我们操作切片可以理解为操作底层的数组。

##### 暴力解法
&emsp;&emsp;第一个for循环遍历切片，第二个for循环将部分数组前移


##### 双指针法
1. 快慢指针
&emsp;&emsp;双指针法（快慢指针法）： 通过一个**快指针**和**慢指针**在**一个for循环**下完成两个for循环的工作。
&emsp;&emsp;在c中erase函数删除元素是将数组中当前位置后面的元素都向前移动一位，其时间复杂度为**O（n）**
&emsp;&emsp;而go中切片的截取（s[:]和append方法），都是通过修改指向底层的数组的指针实现，对于删除头部和尾部元素的时间复杂度为**O（1）**，中间元素的删除则同样为**O（n）**
&emsp;&emsp;完整代码如下，包含学习过程
```go
package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	c := removeElement(nums, val)
	fmt.Println(c)
	fmt.Println(nums)
}

// 暴力解法
/*func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	c := length
	for i, v := range nums {
		if v == val {
			putOff(nums, i, val, length)
			c--
		}
		for nums[i] == val {
			putOff(nums, i, val, length)
			c--
		}

	}
	if nums[0] == val {
		return 0

	}
	return c
}

func putOff(nums []int, i int, val int, length int) {
	if i == length-1 {
		nums[length-1] = val - 1
		return
	}
	nums[i] = nums[i+1]
	putOff(nums, i+1, val, length)
}*/

// 双指针法 这里的解法还是有问题，在对应fast指针更新时，不作更新，只进行fast的右移，可以减少代码量
/*func removeElement(nums []int, val int) int {
	var fast, slow = 0, 0
	length := len(nums)
	for fast < length {
		for nums[fast] == val {
			fast++
			if fast == length {
				return length - fast + slow
			}
		}
		//if nums[slow] == val {
		//
		//}
		nums[slow] = nums[fast]
		slow++
		fast++
	}
	return length - fast + slow
}*/

// 改进双指针，最终版
func removeElement(nums []int, val int) int {
	var fast, slow = 0, 0
	length := len(nums)
	for fast < length {
		if nums[fast] != val { //如果fast对应值不是val，则可以写入
			nums[slow] = nums[fast]
			slow++
		}
		//如果fast对应等于val，则fast移位，不进行更新操作
		fast++

	}
	return slow
}

```
2. 双向指针
```go
//相向双指针法
func removeElement(nums []int, val int) int {
    // 有点像二分查找的左闭右闭区间 所以下面是<=
	left := 0
	right := len(nums) - 1
	for left <= right {
		// 不断寻找左侧的val和右侧的非val 找到时交换位置 目的是将val全覆盖掉
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] == val {
			right--
		}
		//各自找到后开始覆盖 覆盖后继续寻找
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	fmt.Println(nums)
	return left
}
```

#### 有序数组的平方 leetcode977
##### 个人的初次尝试
已经想到双指针法，但是对于循环的边界条件仍然找不明白，对于特殊情况没有适应性，很容易出错
而且在更改的过程中喜欢钻牛角尖，不更改现有的惯性思维，后续需要更正这些缺点。
##### 双指针法
1. 中心法
先找中心位置，从中心开始双指针 从小到大
```go

// 中心位置法，先找到正负交界处的值
func sortedSquares(nums []int) []int {
	//找交界,我们这里找最小的正数
	var i int                //最小正数的序列
	L := len(nums)           //定义nums数组的长度
	result := make([]int, L) //新建结果数组

	for i = 0; nums[i] <= 0; {
		i++
		if i == L {
			break
		}
	}
	// fmt.Println(nums[i])

	//从中心向两边
	// 1.左边到边界
	// 2.右边到边界
	n := 0                      //result的序列
	var nag, pos = i - 1, i     //定义正负两边的指针
	for nag != -1 || pos != L { //循环到边界值，两个都到边界值说明所有值都填完

		if nag == -1 { //非正填完，直接填正即可
			result[n] = nums[pos] * nums[pos]
			pos++
			n++
			continue
		}
		if pos == L { //正填完，直接填非正即可
			result[n] = nums[nag] * nums[nag]
			nag--
			n++
			continue
		}

		if -nums[nag] < nums[pos] { //非正部分小于正，填非正,nag-1
			result[n] = nums[nag] * nums[nag]
			nag--
		} else { //非正部分大于等于正，填正，pos+1
			result[n] = nums[pos] * nums[pos]
			pos++
		}
		n++
	}

	return result
}

```
这里仍旧有瑕疵，比如我们这里在寻找最小正数时，可以将判断n是否等于L加入到判断条件中，而不使用break。后面的填入数字过程中也可以使用多个if else进行判断，不使用continue
2. 两侧双指针
```go
// 两侧双指针法
func sortedSquares(nums []int) []int {
	//由于给定的数组是非递减排序，所以两侧的数平方是最大的，我们从两侧中间寻找最大值反向填充结果数组
	L := len(nums)
	left, right := 0, L-1 //定义左右两侧的指针
	result := make([]int, L)
	index := L - 1 //结果切边填入序列

	for left != right { //当左右指针重合时，说明所有数据已经填充完毕
		leftVal := nums[left] * nums[left]
		rightVal := nums[right] * nums[right]
		if leftVal >= rightVal { //左边值大，填入左边值
			result[index] = leftVal
			left++
		} else {
			result[index] = rightVal
			right--
		}
		index--

	}
	result[0] = nums[left] * nums[right]
	return result
}
```

#### 长度最小的子数组 leetcode209
##### 初次尝试之动态规划 × 超出内存限制
利用如下图所示的思想，但是使用的空间太大了
![Alt text](leetcode209.png)
```go
// 动态规划方法 超出内存限制
func minSubArrayLen(target int, nums []int) int {
	L := len(nums)
	table := make([][]int, L)

	length := 0
	for i := 0; i < L; i++ {
		if nums[i] >= target {
			length = 1
			return length
		}
		table[i] = make([]int, L) //结果表
		table[i][i] = nums[i]     //对角线就是原本数据

		//fmt.Println(table[i])
	}

	//x, y := 0, 1

	//char := 0
outLoop:
	for i := 0; i < L-1; i++ {
		x := 0 //行数都是从0开始
		for y := i + 1; y < L; y++ {
			table[x][y] = table[x][y-1] + table[x+1][y] - table[x+1][y-1]
			if table[x][y] >= target { //如果当前的连续已经超过或等于target
				length = y - x + 1
				break outLoop //第一个大于target的目标即为最短的连续数组
			}
			x++
		}

	}

	for i := 0; i < L; i++ {
		fmt.Println(table[i])
	}

	return length
}
```
##### 滑动窗口(?)我的猪鼻版本
```go
func minSubArrayLen(target int, nums []int) int {
	L := len(nums)
	n := 0 //n用来计算最小的连续数

	left := 0
	right := 0
	add := 0

	for { //第一个for用来找到第一组满足要求的连续数，后续再进行缩小
		right = left + n
		if right == L {
			return 0
		}
		add += nums[right]
		n++
		if add >= target {
			N := n - 1
			for left != L { //第二个for进行缩小
				add -= nums[left]
				left++
				if add >= target {
					n--
					N = n - 1
					continue
				}
				right = left + N
				if right == L {
					break
				}
				add += nums[right]
			}

			return n
		}
	}

}
```

##### 正宗版滑动窗口
```go
func minSubArrayLen(target int, nums []int) int {
	add := 0
	left, right := 0, 0
	L := len(nums)
	length := L + 1 //为了方便判断有无总和超过target

	for right != L {
		add += nums[right]
		for add >= target { //当 当前 窗口内总和大于等于 target 为边界条件；
			length = min(length, right-left+1) //更新最短长度
			add -= nums[left]                  //删除画出的数据
			left++                             //左指针右移
		}
		right++

	}
	if length == L+1 {
		return 0
	}
	return length
}
```


#### 螺旋矩阵 leetcode59
##### 初次尝试，从中心向外
```go
func main() {
	n := 3
	fmt.Println(generateMatrix(n))
}

// 初版，我们从中心点开始
func generateMatrix(n int) [][]int {
	//1.nXn矩阵
	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
		//fmt.Println(table[i]) //确认生成n*n
	}

	//2.确定中心点位置
	//n=1,	(0,0)
	//n=2,	(1,0)
	//n=3,	(1,1)
	//n=4,	(2,1)
	var x, y = 0, 0 //中心点坐标
	x = n / 2
	y = (n - 1) / 2
	fmt.Println("n:", n, "	x:", x, "	y:", y)
	table[x][y] = n * n

	//3.填充

	var char = 0  //都是从横行开始移动, 横竖移动标志位
	var char2 int //到底是向左还是向右
	var char3 int //到底是向上还是向下

	if n%2 == 1 {
		char3 = 1
		char2 = -1
	} else {
		char3 = -1
		char2 = 1
	}
	var c = 1
	for t, i := 1, n*n-1; i >= 0; { //总共填入n*n个数 ,t为当前运行的次数，i为本次填入的数据

		for e := 0; e < 2; e++ {

			tem := (c + 1) / 2 //1,1,2,2,3,3

			if char == 0 { //横着动
				for tem > 0 {
					y = y + char2
					if y == n+1 || y == -1 {
						return table
					}
					table[x][y] = i
					i--
					t++

					tem--
				}
				char2 = -char2
				char = 1
			} else { //竖着动
				for tem > 0 {
					x = x + char3
					if x == n+1 || x == -1 {
						return table
					}
					table[x][y] = i
					i--
					t++
					if t == n*n {
						return table
					}
					tem--
				}
				char3 = -char3
				char = 0

			}
			c++
		}
	}

	return table
}

```

##### 模拟方式
```go
func generateMatrix(n int) [][]int {

	//1.nXn矩阵
	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
		//fmt.Println(table[i]) //确认生成n*n
	}

	i := 1 //需要填入的数据

	c := n - 1

	layerNums := 0            //第几层

	for ; c >= 0; c = c - 2 { //记每次循环每层每边中间间隔的个数，c=0时为最后中心只有一个数的那层
	//n=2 2X2 C:1
	//n=3 3X3 C:2 -> 0
	/*
	1 2 3 
	4 5 6
	7 8 9
 	*/
	//n=4 4X4 C:3 -> 1 
	/* 
	1  2  3  4
	12 13 14 5
	11 16 15 6
	10  9  8 7
	*/
	
	x, y := layerNums, layerNums //记每层模拟的起始点
		table[x][y] = i
		for b := 0; b < c; b++ {//先向右横
			table[x][y] = i
			y++
			i++
		}
		for b := 0; b < c; b++ {//向下竖
			table[x][y] = i
			x++
			i++
		}
		for b := 0; b < c; b++ {//向左横
			table[x][y] = i
			y--
			i++
		}
		for b := 0; b < c; b++ {//向上竖
			table[x][y] = i
			x--
			i++
		}
		//由于我们写入实际在下次循环才写入，所以最后一位的原点实际并未覆盖


		layerNums++//层数计数加一
	}

	return table
}

```

### 链表
更多有关于go链表的内容可以见这篇文章[链表的创建和读取](https://blog.csdn.net/kingsill/article/details/134968248?spm=1001.2014.3001.5501)
#### 链表基础知识部分
什么是**链表**，链表是一种通过**指针**串联在一起的**线性结构**，每一个节点由两部分组成，一个是**数据域**一个是**指针域**（存放指向下一个节点的指针），最后一个节点的指针域指向null（空指针的意思）。

链表的入口节点称为链表的**头结点**也就是`head`

- 链表类型
	- 单链表
	![Alt text](singleList.png)
	如上
	- 双链表：	 
	每一个节点有两个指针域，一个指向下一个节点，一个指向上一个节点。
	![Alt text](doubleList.png)
	- 循环链表
	链表首尾相连
	![Alt text](circularList.png)

##### 链表的存储方式
数组是在内存中是连续分布的，但是链表在内存中可不是连续分布的。

##### 链表的定义
go语言：单链表
```go
type ListNode struct {
      Val int 		//数据域
      Next *ListNode//指针域
}
```

##### 链表的操作
1. 节点删除
只需要将删除的节点后面的节点连接到前面的节点即可
由于go有自己的内存回收机制，删除的节点不用再手动释放
![Alt text](deleteNode.png)
2. 添加节点
将要添加的节点指针域指向后面的节点，前面的指针域替换为要添加节点的
![Alt text](addNode.png)

##### 性能分析
|时间复杂度|插入/删除|查询|适用场景|
|:---:|---|---|---|
|数组|O(n)|O(1)|数据量固定，频繁查询，较少增删|
|链表|O(1)|O(n)|数据量不固定，频繁增删，较少查询|

数组在定义的时候，长度就是固定的，如果想改动数组的长度，就需要重新定义一个新的数组。

链表的长度可以是不固定的，并且可以动态增删， 适合数据量不固定，频繁增删，较少查询的场景。

更多有关于go链表的内容可以见这篇文章[链表的创建和读取](https://blog.csdn.net/kingsill/article/details/134968248?spm=1001.2014.3001.5501)


#### 移除链表元素 leetcode203

##### 初版方法 迭代 虚拟头节点
使用迭代解决 ，虚拟头节点
没有使用给出的链表，选择重建了一个链表，时间、空间复杂度都增加
```go
func removeElements(head *ListNode, val int) *ListNode {
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

	return targetList.Next
}

```

##### 改进 迭代 虚拟头节点
```go
// 新版本 迭代 虚拟头节点
func removeElements(head *ListNode, val int) *ListNode {
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

```

###### 使用迭代 直接使用原链表 不使用头节点
```go
func removeElements(head *ListNode, val int) *ListNode {

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

}

```

###### 递归方法 
首先是建立链表的思路
我选择寻求gpt的帮助 ：）

构建递归的思路通常遵循以下几个步骤：
- **确定递归结束条件**：这是最重要的一步，递归必须有一个明确的结束条件，否则会导致无限递归和栈溢出。结束条件通常是当问题规模缩减到最小或达到某个特定状态时。

- **找到问题的递归式**：确定如何将问题分解成更小的子问题。递归的核心思想是将大问题分解成结构或逻辑上类似的小问题。

- **调用自身解决子问题**：在函数中调用自身（递归调用）来解决这些子问题。这些子问题的解决方式应与原问题相同。

- **整合子问题的解决结果**：递归调用后，通常需要整合或处理这些子问题的结果来得到最终答案。

- **优化（可选）**：对于某些递归问题，可能会出现重复计算的情况，可以通过技巧如记忆化（memoization）来优化性能。

那么我们将结合leetcode203进行上述思路进行解题


要使用递归方法解决删除链表中所有满足 Node.val == val 的节点问题，可以遵循以下递归思路：

- **递归结束条件**：当当前节点为 null 时，表示已经到达链表末尾，递归结束。在这种情况下，直接返回 null。

- **递归式**：
	检查当前节点（假设为 head）的值是否等于 val。
	如果是，我们需要删除这个节点，这意味着需要将其父节点（递归中的上一个节点）连接到它的下一个节点。
	但在递归中，我们没有直接的方式来访问父节点，因此我们通过返回下一个节点（head.next）来实现这一点。

- **调用自身**：对于每个节点，调用递归函数处理它的下一个节点（即，调用自身传递 head.next 和 val）。

- **整合子问题的解决结果**：如果当前节点的值不等于 val，那么我们需要保留这个节点，并将其 next 指针指向递归调用的结果。否则，直接返回递归调用的结果，即 head.next 的递归处理结果。

- **返回新的头节点**：递归的最终结果是新的链表头节点，它可能与原始链表的头节点不同，尤其是当原始头节点的值等于 val 时。

```go
// 递归
func removeElements(head *ListNode, val int) *ListNode {

	//放在递归开头的为 结束条件 ，其应该return一个结果，结束本次递归
	//即本次递归传入的 head为空，不再进行后续递归，将尾节点设置为空
	if head != nil {
		return head
	}

	//调用自身
	//修改子节点
	head.Next = removeElements(head.Next, val) //我们也可以通过这个式子理解 下面的返回 返回到.next

	//递归式-整合子问题的解决成果
	//返回到父节点
	if head.Val == val {
		return head.Next //如果当前节点的val等于val则舍弃该节点
	}
	return head
}

```
