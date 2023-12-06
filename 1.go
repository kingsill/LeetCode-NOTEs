package main

import "fmt"

var nums = []int{2, 7, 11, 15}
var target = 9

func main() {
	twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	// var result map[int]int
	for i, v := range nums {

		fmt.Printf("v: %v\n", v)
		fmt.Printf("i: %v\n", i)
	}
	return nil

}
