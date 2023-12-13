package main

import "fmt"

var nums1 = []int{}
var nums2 = []int{1, 2}

func main() {

	f := findMedianSortedArrays(nums1, nums2)
	fmt.Println(f)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var N = len(nums1)
	var M = len(nums2)
	z := M + N //两个数组一共有多少位数
	fmt.Println("z:", z)
	if N == 0 {
		return SingleOrder(nums2)
	}
	if M == 0 {
		return SingleOrder(nums1)
	}
	if z%2 == 1 { //奇数
		I := (M + N + 1) / 2             //I用来指代需要排序到第几位数
		i := OrderSlice(nums1, nums2, I) //i指代获得的排序切片
		var mid float64 = float64(i[I-1])
		return mid
	} else { //偶数
		I := (M+N)/2 + 1
		i := OrderSlice(nums1, nums2, I)
		var mid float64 = (float64(i[(M+N)/2-1]) + float64(i[I-1])) / 2
		return mid
	}
}

// OrderSlice 传入的I是要排序到第几位
func OrderSlice(nums1 []int, nums2 []int, I int) []int {
	var order []int //用于存储切片的实例

	n1, n2 := 0, 0 //记录nums1、nums2读取到第几个

	for i := 0; i < I; i++ { //一共排列I个数

		if nums1[n1] <= nums2[n2] { //n1<=n2,则排列n1中的数
			order = append(order, nums1[n1])
			n1++
			if n1 == len(nums1) { //当n1排完
				order = append(order, nums2[n2:]...)
				break
			}

		} else {
			order = append(order, nums2[n2])
			n2++
			if n2 == len(nums2) { //当n1排完
				order = append(order, nums1[n1:]...)
				break
			}

		}
	}
	fmt.Printf("order: %v\n", order)
	return order
}

func SingleOrder(nums []int) float64 {
	i := len(nums)

	if i%2 == 1 { //奇数
		return float64(nums[(i+1)/2-1])
	} else { //偶数
		return (float64(nums[i/2-1]) + float64(nums[i/2+1-1])) / 2
	}
}
