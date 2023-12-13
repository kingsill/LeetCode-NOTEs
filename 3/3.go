package main

import (
	"fmt"
)

func main() {
	s := "zwnigfunjwz"
	c := lengthOfLongestSubstring(s)
	fmt.Println(c)
}

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
