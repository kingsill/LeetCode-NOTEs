package main

import (
	"fmt"
)

func main() {

	c := isHappy(19)
	fmt.Println(c)
}

/*
func isHappy(n int) bool {
	result := 0
	temp := 0
	record := make(map[int]int)

	for result != 1 {

		for ; n != 0; n = n / 10 {

			c := n % 10
			temp += c * c
			result = temp
		}
		if _, ok := record[result]; ok {
			return false
		}
		record[result] = 1
		n = result
		temp = 0
	}
	return true

}
*/

// 快慢指针法
func isHappy(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
