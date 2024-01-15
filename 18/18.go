package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-2, -1, -1, 1, 1, 2, 2}
	fmt.Println(fourSum(nums, 0))
}

// 自我救赎的双指针法
func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	L := len(nums)
	if L < 4 {
		return nil
	}

	record := [][]int{}

	//第一个数
	for i, val := range nums[:L-3] {
		if i > 0 && val == nums[i-1] {
			continue
		}

		//第二个数
		for i2 := i + 1; i2 < L-2; i2++ {
			val2 := nums[i2]
			//左指针为第二个数，右指针为第三个数。左指针每次从第一个数后一个开始，右指针则从最末尾开始
			left, right := i2+1, L-1

			//第一个数的去重 ，如{-1 -1 0 1}，[0,2,3]和[1,2,3]都满足要求，但是其数组都为[-1,0,1]应该去除之一。所以说当检测当本次循环与上次一致的话，直接跳过
			if i2 > i+1 && val2 == nums[i2-1] {
				continue
			}

			//循环的结束条件为左指针与右指针重合，即第二个数和第三个数序号一致
			for left < right {
				if val+val2+nums[left]+nums[right] == target {
					record = append(record, []int{val, val2, nums[left], nums[right]})
					left++
					right--

					//对第二个数进行去重
					for nums[left] == nums[left-1] && left < L-2 {
						left++
					}

					//对第三个数进行去重
					for nums[right] == nums[right+1] && right > 2 {
						right--
					}
				}
				if val+val2+nums[left]+nums[right] > target {
					right--
				}
				if val+val2+nums[left]+nums[right] < target {
					left++
				}
			}

		}

	}

	return record
}
