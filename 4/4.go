package main

import (
	"fmt"
)

var nums1 = []int{1, 6}
var nums2 = []int{2, 3, 4, 5, 7, 8, 9, 10}

func main() {

	f := findMedianSortedArrays(nums1, nums2)

	fmt.Println(f)

}

/*func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
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
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	z := len(nums1) + len(nums2)
	if z%2 == 1 { //奇数
		return float64(Get(nums1, nums2, (z+1)/2))
	} else { //偶数
		return (float64(Get(nums1, nums2, z/2) + Get(nums1, nums2, z/2+1))) / 2
	}
}

func Get(nums1 []int, nums2 []int, k int) int {
	var Rindex1, Rindex2 = 0, 0
	offset1, offset2 := 0, 0
	//Vindex1, Vindex2 := 0, 0
	var flow int = 0

	for {
		if Rindex1 == len(nums1) {
			return nums2[k+Rindex2-1]
		}
		if Rindex2 == len(nums2) {
			return nums1[k+Rindex1-1]
		}

		if k == 1 {

			return min(nums1[Rindex1], nums2[Rindex2])
		}
		point1 := k/2 - 1 + Rindex1
		point2 := k/2 - 1 + Rindex2

		//判断有没有溢出
		if k/2-1+Rindex1 > len(nums1)-1 {
			point1 = len(nums1) - 1
			flow = 1
		}
		if k/2-1+Rindex2 > len(nums2)-1 {
			point2 = len(nums2) - 1
			flow = 2
		}
		if flow == 1 || flow == 2 {
			if flow == 1 { //如果时数组1超限
				if nums1[point1] <= nums2[point2] {
					offset1 = len(nums1)
					k = k - len(nums1)
				} else {
					offset2 = k / 2
					k -= k / 2
				}

			}
			if flow == 2 { //如果是数组2超限
				if nums1[point1] <= nums2[point2] {
					offset1 = k / 2
					k -= k / 2
				} else {
					offset2 = len(nums2)
					k = k - len(nums2)
				}

			}

		} else {
			if nums1[point1] <= nums2[point2] {
				offset1 = k / 2
			} else {
				offset2 = k / 2
			}

			k -= k / 2
		}

		Rindex1 += offset1
		Rindex2 += offset2
		offset1, offset2, flow = 0, 0, 0
	}

}

/*func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
*/
