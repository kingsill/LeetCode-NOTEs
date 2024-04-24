package main

import (
	"fmt"
)

func main() {
	nums := []int{-6, -10, -7, -1, -9, 9, -8, -4, 10, -5, 2, 9, 0, -7, 7, 4, -2, -10, 8, 7}
	fmt.Println(maxSlidingWindow(nums, 7))

}

func maxSlidingWindow(nums []int, k int) []int {
	//define a stack
	stack := make([]int, k)

	//define the result
	L := len(nums)
	result := make([]int, L)

	for i, v := range nums {
		if i >= k {
			stack = pop(stack, v)
			stack = push(stack, nums[i-k])
		} else {
			if i == 0 {
				stack[0] = nums[0]
			} else {
				stack = pop(stack, v)
			}

		}
		result[i] = getMax(stack)
	}
	return result[k-1:]
}

func pop(stack []int, v int) []int {
	// Remove elements from the end of the stack that are smaller than v
	for len(stack) > 0 && v > stack[len(stack)-1] {
		stack = stack[:len(stack)-1]
	}
	// Append v to the stack
	stack = append(stack, v)
	return stack
}

func push(stack []int, v int) []int {
	if v == stack[0] {
		stack = stack[1:]
	}
	return stack
}

func getMax(stack []int) int {
	return stack[0]
}
