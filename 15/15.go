package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{1, -1, 0, -1, 1}
	c := threeSum(nums)
	fmt.Printf("c: %v\n", c)

}

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
						intT = fmt.Sprintf(intT, string(rune(v)))
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
