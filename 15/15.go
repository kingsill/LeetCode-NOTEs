package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	c := threeSum(nums)
	fmt.Printf("c: %v\n", c)

}

// 得到结果后再去重 失败,超出输出限制
func threeSum(nums []int) [][]int {
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
}

// 双指针法
/*func threeSum(nums []int) [][]int {

	slices.Sort(nums)
	L := len(nums)

	result := [][]int{}

	for left := 0; left < L; left++ {
		for middle := left + 1; middle < L; middle++ {
			for right := middle + 1; right < L; right++ {
				if nums[left]+nums[middle]+nums[right] == 0 {
					result = append(result, []int{nums[left], nums[middle], nums[right]})
				}

			}
		}
	}

	fmt.Println(nums)

	return nil
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
