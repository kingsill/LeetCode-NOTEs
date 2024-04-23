package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
