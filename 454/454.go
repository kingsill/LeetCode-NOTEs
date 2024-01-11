package main

func main() {

}

// 使用两遍for循环
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	index := 0

	record := map[int]int{}

	//首先对num3、nums4进行遍历，记录数据	考虑到可能每组有重复数字，每遍历到一次就+1
	for _, i1 := range nums3 {
		for _, i2 := range nums4 {
			record[i1+i2]++
		}

	}

	//然后再进行nums1、nums2的遍历
	for _, i1 := range nums1 {
		for _, i2 := range nums2 {
			//record对应数字是几，证明能得到这个数字的组合有几个，这里就加上对应的组合数
			if v, ok := record[-i1-i2]; ok {
				index = index + v
			}
		}

	}
	return index
}
