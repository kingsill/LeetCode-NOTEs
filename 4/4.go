package main

import "fmt"

var nums1 = []int{1, 3}
var nums2 = []int{2, 4}

var N = len(nums1)
var M = len(nums2)

func main() {

	f := findMedianSortedArrays(nums1, nums2)
	fmt.Println(f)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	z := M + N //两个数组一共有多少位数

	if z%2 == 1 { //奇数
		I := (M + N + 1)
		i := OrderSlice(nums1, nums2, I)
		var mid float64 = float64(i[I])
		return mid
	} else { //偶数
		I := (M+N)/2 + 1
		i := OrderSlice(nums1, nums2, I)
		var mid float64 = (float64(i[(M+N)/2-1]) + float64(i[I-1])) / 2
		return mid
	}
}

// 传入的I是要排序到第几位
func OrderSlice(nums1 []int, nums2 []int, I int) []int {
	order := []int{} //用于存储切片的实例
	fmt.Printf("nums1: %v\n", nums1)
	n1, n2 := 0, 0 //记录nums1、nums2读取到第几个

	for i := 0; i < I; i++ { //一共排列I个数
		if nums1[n1] <= nums2[n2] { //n1<=n2,则排列n1中的数

			order = append(order, nums1[n1])
			n1++
		} else {
			order = append(order, nums2[n2])
			n2++
		}
	}
	fmt.Printf("order: %v\n", order)
	return order
}
