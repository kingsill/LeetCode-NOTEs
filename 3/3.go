package main

import (
	"fmt"
)

func main() {
	s := "bbbbbbbb"
	c := lengthOfLongestSubstring(s)
	fmt.Println(c)
}

//我的个人第一解版本
/*
	func lengthOfLongestSubstring(s string) int {
	//设置为每次保存最大个数的窗口
	//type Win map[string]int

	//每次循环的窗口
	n := make(map[string]int)

	//保存每次循环得到的最大值
	var Max int

	//记录每个窗口中最多的字母数
	i := 0
	var light int = 1     //用作窗口的左侧指针
	for c, v := range s { //c+1用来记录找到第几个位置
		//fmt.Println(string(v))
		fmt.Println(n)
		fmt.Println("c+1(n):", c+1)

		_, ok := n[string(v)]

		//fmt.Println(ok)

		if ok { //如果有对应的字母，说明有重复,1.清空当前窗口 2.比较临时最长长度和记录的最长长度，保留较大值 3.清空临时长度计数
			//先进行计算，再进行下一个窗口的布置
			if n[string(v)] >= light { //如果这个重复的字母在light的右边，则需要用这个字母的位置进行最长长度的jisuan
				fmt.Println("一")
				i = c + 1 - n[string(v)]
			} else { //否则用light计算即可
				fmt.Println("二")
				i = c + 1 - light
			}

			light = max(light, n[string(v)]) //更新light的位置
			n[string(v)] = c + 1             //更新重复字母最后的位置
			fmt.Println("light:", light)

			fmt.Println("i:", i)

			//n = map[string]int{string(v): i} //保留当前重复的一位，从计数1开始计数
		} else { //如果没有对应的字母1.最长长度计数+1 2.将临时最长长度长度保存

			n[string(v)] = c + 1
			//MaxTemp = i
			i++
		}

		//在遍历的过程中不断更新最长长度，防止最后一段字母是最长的
		if i >= Max {
			Max = i
			fmt.Println("MAX:", Max)
		}
	}
	return Max
	}
*/

// 官方解法
/*func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}*/

// 梅开二度，第二次进行尝试 抄袭官方版本
func lengthOfLongestSubstring(s string) int {
	//用于记录遍历到的字符是否出现过
	Record := map[int32]bool{}
	//Record := [128]bool{} //由于asc2字符一共有128个，这里就用128个数字代表所有的字符，range得到的内容是int32格式的，正好方便写入

	var left int = 0

	var Max int = 1

	for right, v := range s { //遍历s

		//如果有重复字符，左指针右移，窗口缩小直到没有重复字符
		for Record[v] == true {
			Record[int32(s[left])] = false
			delete(Record, int32(s[left]))
			left++ //左指针右移

		}

		//如果没有重复字符
		Record[v] = true             //将出现的字符标记为true
		Max = max(Max, right-left+1) //计算最大无重复字符长度

	}

	return Max
}
