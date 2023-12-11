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
 并未自行接出来
