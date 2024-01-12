package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-2, -1, -1, 1, 1, 2, 2}
	fmt.Println(fourSum(nums, 0))
}

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	L := len(nums)

	fourth := L - 1

	record := [][]int{}

	for first := 0; first < L; first++ {
		if first > 1 && nums[first] == nums[first-1] && nums[first] == nums[first-2] {
			continue
		}
		for second := first + 1; second < L; second++ {
			if second > 1 && nums[second] == nums[second-1] && nums[second] == nums[second-2] {
				continue
			}
			for third := second + 1; third < L; third++ {
				if third > 2 && nums[third] == nums[third-1] && nums[third] == nums[third-2] {
					continue
				}
				for nums[first]+nums[second]+nums[third]+nums[fourth] > target && third < fourth {
					fourth--
				}
				if fourth == third {
					break
				}
				if nums[first]+nums[second]+nums[third]+nums[fourth] == target {
					record = append(record, []int{nums[first], nums[second], nums[third], nums[fourth]})
				}

			}
			if fourth == second+1 {
				break
			}
		}
	}
	return record
}
