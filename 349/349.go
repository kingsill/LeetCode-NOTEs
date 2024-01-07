package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 3, 4}
	fmt.Println(intersection(nums1, nums2))

}

// 最简易的方式
/*func intersection(nums1 []int, nums2 []int) []int {
	table := make(map[int]int)

	result := []int{}

	for _, i2 := range nums1 {
		table[i2] = 1
	}
	for _, i2 := range nums2 {
		if table[i2] == 1 {
			table[i2] = 0
			result = append(result, i2)
		}

	}
	return result
}
*/

// 排序+双指针
func intersection(nums1 []int, nums2 []int) []int {
	index1, index2 := 0, 0

	sort.Ints(nums1)
	sort.Ints(nums2)

	L1 := len(nums1)
	L2 := len(nums2)

	result := []int{}

	for index1 < L1 {

		//当出现不是等于的情况时直接for循环
		for nums1[index1] < nums2[index2] {
			index1++
			if index1 >= L1 {
				return result
			}
		}

		for nums1[index1] > nums2[index2] {
			index2++
			if index2 >= L2 {
				return result
			}
		}

		if nums1[index1] == nums2[index2] {
			result = append(result, nums1[index1])
			for nums1[index1] == nums2[index2] {
				index2++
				if index2 >= L2 {
					return result
				}
			}
		}
	}
	return result

}
