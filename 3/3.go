package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	c := lengthOfLongestSubstring(s)
	fmt.Println(c)
}

func lengthOfLongestSubstring(s string) int {
	//设置为每次保存最大个数的窗口
	//type Win map[string]int

	//每次循环的窗口
	n := make(map[string]int)

	//保存每次循环得到的最大值
	var MaxTemp, Max int

	//记录每个串口中最多的字母数
	i := 0
	for _, v := range s {
		fmt.Println(string(v))
		fmt.Println(n)
		_, ok := n[string(v)]
		fmt.Println(ok)
		if ok { //如果有对应的字母，说明有重复,1.清空当前窗口 2.比较临时最长长度和记录的最长长度，保留较大值 3.清空临时长度计数
			n = map[string]int{}
			if MaxTemp >= Max {
				Max = MaxTemp
			}

			i = 0
		} else { //如果没有对应的字母1.最长长度计数+1 2.将临时最长长度长度保存
			i++
			n[string(v)] = i
			MaxTemp = i
		}
	}
	return Max
}
