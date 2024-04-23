package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2}
	i := sortedSquares(nums)
	fmt.Printf("i: %v\n", i)
}

//初次的失败
/* func sortedSquares(nums []int) []int {
	L := len(nums)
	n := 0
	result := make([]int, L)
	p := 0
	c := 0
	for i, v := range nums {
		if v <= 0 {
			n++

			result[L-1-i] = v * v //将小于等于0的部分倒序插入到result中
		} else {
			fmt.Printf("n: %v\n", n)

			fmt.Println(v * v)
			for ; v*v > result[L-n+c]; p++ {
				fmt.Printf("p: %v\n", p)
				result[p] = result[L-n+c]
				c++
				fmt.Printf("c: %v\n", c)
				if c == n {
					result[p+1] = v * v
					return result
				}
			}

			result[p] = v * v
			fmt.Printf("p: %v\n", p)
			p++
		}
	}
	return result
}
*/

//深入了解部分后的尝试
//  1.先找中心位置，从中心开始双指针 从小到大
//  2.从两边开始双指针 从大到小

// 中心位置法，先找到正负交界处的值
/* func sortedSquares(nums []int) []int {
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
} */

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

// 重来版本
func sortedSquares(nums []int) []int {
	low, high, index := 0, len(nums)-1, len(nums)-1
	result := make([]int, high+1)
	for low <= high {
		left := nums[low] * nums[low]
		right := nums[high] * nums[high]

		if left > right {
			result[index] = left
			low++
		} else {
			result[index] = right
			high--
		}

		index--
	}
	return result
}
