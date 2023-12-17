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


