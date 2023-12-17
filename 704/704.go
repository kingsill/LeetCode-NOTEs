package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	n := search(nums, 8)
	fmt.Println(n)
}

//赖子玩法，使用hashmap进行解决
/*func search(nums []int, target int) int {
	Abook := map[int]int{}
	for i, num := range nums {
		Abook[num] = i
	}
	number, ok := Abook[target]
	if ok {
		return number
	}
	return -1
}
*/

// 使用二分查找进行查询 个人版本，对区间的定义理解不够
// 因此需要额外定义多种情况
/*func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	//如果该目标值直接大于或者小于nums中所有数，直接返回-1
	if target > nums[right] || target < nums[left] {
		return -1
	}

	//为了避免与后面查询不到的情况避开，我们这里直接设定两个值时的情况
	if nums[0] == target {
		return 0
	}
	if nums[right] == target {
		return right
	}

	for {
		mid := (left + right) / 2 //偶数为中间往前一个数 ； 奇数即为中间数

		//如果寻找到目标值。返回
		if nums[mid] == target {
			return mid
		}

		//如果没有查询到目标值
		if math.Abs(float64(left-right)) == 1 {
			return -1
		}

		if nums[mid] < target { //目标大于中间值，往后查找
			left = mid
		} else { //nums[mid] > target//目标小于中间值，往前查询
			right = mid
		}

	}

}*/

// 学习版本二分查找 闭区间
/*func search(nums []int, target int) int {
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
}*/

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
