package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	target := 11
	c := minSubArrayLen(target, nums)
	fmt.Println("c:", c)
}

// 动态规划方法 超出内存限制
/*func minSubArrayLen(target int, nums []int) int {
	L := len(nums)
	table := make([][]int, L)

	length := 0
	for i := 0; i < L; i++ {
		if nums[i] >= target {
			length = 1
			return length
		}
		table[i] = make([]int, L) //结果表
		table[i][i] = nums[i]     //对角线就是原本数据

		//fmt.Println(table[i])
	}

	//x, y := 0, 1

	//char := 0
outLoop:
	for i := 0; i < L-1; i++ {
		x := 0 //行数都是从0开始
		for y := i + 1; y < L; y++ {
			table[x][y] = table[x][y-1] + table[x+1][y] - table[x+1][y-1]
			if table[x][y] >= target { //如果当前的连续已经超过或等于target
				length = y - x + 1
				break outLoop //第一个大于target的目标即为最短的连续数组
			}
			x++
		}

	}

	for i := 0; i < L; i++ {
		fmt.Println(table[i])
	}

	return length
}
*/
//滑动窗口(?)
/*func minSubArrayLen(target int, nums []int) int {
	L := len(nums)
	n := 0 //n用来计算最小的连续数

	left := 0
	right := 0
	add := 0

	for { //第一个for用来找到第一组满足要求的连续数，后续再进行缩小
		right = left + n
		if right == L {
			return 0
		}
		add += nums[right]

		n++
		if add >= target {
			break
		}
	}
	N := n - 1
	for left != L { //第二个for进行缩小
		add -= nums[left]
		left++
		if add >= target {
			n--
			N = n - 1
			continue
		}
		right = left + N
		if right == L {
			break
		}
		add += nums[right]
	}

	return n
}
*/

/*func minSubArrayLen(target int, nums []int) int {
	L := len(nums)
	n := 0 //n用来计算最小的连续数

	left := 0
	right := 0
	add := 0

	for { //第一个for用来找到第一组满足要求的连续数，后续再进行缩小
		right = left + n
		if right == L {
			return 0
		}
		add += nums[right]
		n++
		if add >= target {
			N := n - 1
			for left != L { //第二个for进行缩小
				add -= nums[left]
				left++
				if add >= target {
					n--
					N = n - 1
					continue
				}
				right = left + N
				if right == L {
					break
				}
				add += nums[right]
			}

			return n
		}
	}
}*/

func minSubArrayLen(target int, nums []int) int {
	i := 0          //左侧指针的位置
	l := len(nums)  // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况

	for j := 0; j < l; j++ { //j为右侧指针，当右指针到末位时，到达边界
		sum += nums[j]      //累计求和
		for sum >= target { //边界条件为sum是否大于等于target
			subLength := j - i + 1  //求满足条件时的子数组长度，这里为临时
			if subLength < result { //比较此次的临时长度和保存的历史最短，选择最短则进行保存
				result = subLength
			}
			sum -= nums[i] //将窗口向右滑动，将划出的数据删除
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}
