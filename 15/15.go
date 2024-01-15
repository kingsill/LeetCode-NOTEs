package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-2, 0, 1, 1, 2}
	c := threeSum(nums)
	fmt.Printf("c: %v\n", c)

}

// 得到结果后再去重 失败,超出输出限制
/*func threeSum(nums []int) [][]int {
	L := len(nums)

	var intT string

	result := [][]int{}

	N := map[int]int{}
	G := map[string]int{}

	table := make([][]int, L)

	for i, _ := range table {
		table[i] = make([]int, L)
	}

	for i, v := range nums {
		N[v] = i
	}

	for i := 0; i < L-1; i++ {
		for c := i + 1; c < L; c++ {
			table[i][c] = nums[i] + nums[c]
			if index, ok := N[-table[i][c]]; ok {
				if index != i && index != c {
					order := []int{nums[i], nums[c], nums[index]}

					slices.Sort(order)

					for _, v := range order {
						intT = fmt.Sprintf(intT, v)
						fmt.Println(intT)
					}

					if _, cont := G[intT]; !cont {
						result = append(result, []int{nums[i], nums[c], nums[index]})
					}
					G[intT] = 1
					intT = ""
				}
			}

		}
	}
	// fmt.Print(table)
	return result
}*/

// 官方解题
/*func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 first
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// third 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 second
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 second 的指针在 third 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 second 后续的增加
			// 就不会有满足 first+second+third=0 并且 second<third 的 third 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
*/

// 自我救赎的双指针法
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	L := len(nums)
	record := [][]int{}

	//第一个数
	for i, i2 := range nums[:L-2] {

		//左指针为第二个数，右指针为第三个数。左指针每次从第一个数后一个开始，右指针则从最末尾开始
		left, right := i+1, L-1

		//第一个数的去重 ，如{-1 -1 0 1}，[0,2,3]和[1,2,3]都满足要求，但是其数组都为[-1,0,1]应该去除之一。所以说当检测当本次循环与上次一致的话，直接跳过
		if i > 0 && i2 == nums[i-1] {
			continue
		}

		//循环的结束条件为左指针与右指针重合，即第二个数和第三个数序号一致
		for left < right {
			if i2+nums[left]+nums[right] == 0 {
				record = append(record, []int{i2, nums[left], nums[right]})
				left++
				right--

				//对第二个数进行去重
				for nums[left] == nums[left-1] && left < L-1 {
					left++
				}

				//对第三个数进行去重
				for nums[right] == nums[right+1] && right > 1 {
					right--
				}
			}
			if i2+nums[left]+nums[right] > 0 {
				right--
			}
			if i2+nums[left]+nums[right] < 0 {
				left++
			}
		}

	}
	return record
}
