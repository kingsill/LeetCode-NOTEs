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
/*func removeElement(nums []int, val int) int {
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
}*/

// 双向指针版
// 相向双指针法
/*func removeElement(nums []int, val int) int {
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
*/

// 重来 轻松拿下
func removeElement(nums []int, val int) int {
	//left指针用来赋值，right用来寻值
	left := 0
	//right到头时结束

	for _, rightVal := range nums {
		if rightVal != val {
			nums[left] = rightVal
			left++
		}
	}
	return left
}
