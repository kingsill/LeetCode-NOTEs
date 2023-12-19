package main

import "fmt"

func main() {
	nums := []int{1}
	i := sortedSquares(nums)
	fmt.Printf("i: %v\n", i)
}

func sortedSquares(nums []int) []int {
	L := len(nums)
	n := 0
	result := make([]int, L)
	p := 0
	c := 0
	for i, v := range nums {
		if v <= 0 {
			n++

			result[L-1-i] = v * v //将小于等于0的部分倒序插入到result中
		} else {
			fmt.Printf("n: %v\n", n)

			fmt.Println(v * v)
			for ; v*v > result[L-n+c]; p++ {
				fmt.Printf("p: %v\n", p)
				result[p] = result[L-n+c]
				c++
				fmt.Printf("c: %v\n", c)
				if c == n {
					result[p+1] = v * v
					return result
				}
			}

			result[p] = v * v
			fmt.Printf("p: %v\n", p)
			p++
		}
	}
	return result
}
