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
