package main

import (
	"fmt"
)

func main() {
	haystack := "mississippi"
	needle := "issip"
	num := strStr(haystack, needle)
	fmt.Println(num)

	temp := initNext(needle)
	fmt.Println(temp)

}

// 算是暴力算法，将needle与每一个开头字母与needle相同的子字符串进行比较
/*func strStr(haystack string, needle string) int {
	L := len(needle)
	Cap := len(haystack)

	H := []byte(haystack)
	N := []byte(needle)

	for i, val := range H {

		if val == N[0] {
			if i+L <= Cap && haystack[i:i+L] == needle {
				return i
			}
		}
	}

	return -1
}*/

// kmp算法，用空间换时间
func strStr(haystack string, needle string) int {
	//获取next数组
	next := initNext(needle)

	//主串长度
	L := len(haystack)

	//目标匹配长度，即needle的长度
	target := len(needle)

	//匹配字符串中指针位置
	j := 0

	//i为主串中指针的位置
	for i := 0; i < L; {

		// 如果匹配上了
		if haystack[i] == needle[j] {
			if j == target-1 {
				return i - target + 1
			}
			j++
			i++
		} else { //如果没匹配上

			//跟计算next数组有异曲同工之妙
			if j > 0 {
				j = next[j-1]
			} else {
				i++
			}

		}
	}

	return -1
}

// 此函数用来初始化next数组
func initNext(needle string) []int {
	//后缀中末尾  abc中c
	i := 1

	//前缀中末尾;同时在这里也有着记录最长公共串长度的作用，二者本质是一样的
	j := 0

	L := len(needle)

	//初始化next数组；next[0]默认为0，因为对于一个字母我们不认为其具有前后缀，后续也不会再对next[0]进行赋值
	next := make([]int, L)

	//求next数组过程中，我们的i不回退，采用类似于动态规划的思想，也是我们这里的循环条件
	for i < L {

		//如果前后缀匹配
		if needle[i] == needle[j] {

			//前缀末尾向后移一位，同时代表长度+1
			j++

			//当前后缀末尾所在位置的最长子串即为j
			//最长子串是有基础的，如果next[2]=2，那么next[3]的可能性为3或者0,这里是为3的情况
			next[i] = j

			//后缀末尾向后移一位
			i++

		} else { //如果前后缀不匹配

			//当j>0，说明仍旧处于回退的过程
			if j > 0 {
				j = next[j-1]
			} else { //如果j=0，并且前后缀依旧不匹配，则长度计数应该重新从0开始

				//这里是为0的情况
				next[i] = j

				//后缀末尾向后移
				i++
			}
		}
	}
	//返回next数组
	return next
}
