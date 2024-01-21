package main

import (
	"fmt"
)

func main() {
	fmt.Println(repeatedSubstringPattern("ababab"))
}

//思想滑坡----------------------------
/*func repeatedSubstringPattern(s string) bool {
	//首先统计各个字母出现的次数---------------------
	record := make([]int, 26)
	for _, i2 := range s {
		record[i2-'a']++
	}

	//如果有最小公因数k，初步判断达成，否则return false--------------------
	k := 10 * 10 * 10 * 10
	for _, val := range record {
		if val != 0 {
			k = min(k, val)
			if val%k != 0 {
				return false
			}
		}
	}

	//继续判断，将字符串s平分为k份，比较是否都一致，一致则return ture，否则为flase------------------
	single := len(s) / k
	if len(s)%k != 0 || k == 1 {
		return false
	}
	n := k
	for ; n > 1; n-- {
		if s[(n-1)*single:n*single] != s[:single] {
			return false
		}
	}
	return true
}*/

// 暴力枚举------------------------------------
/*func repeatedSubstringPattern(s string) bool {

	L := len(s)

	record := false

	for i := 1; i < L/2+1; i++ {
		if L%i != 0 {
			continue
		}
		num := L / i

		for ; num > 0; num-- {
			if s[:i] != s[(num-1)*i:num*i] {
				record = false
				break
			}
			record = true
		}
		if record == true {
			return record
		}
	}
	return record

}*/

// 移动匹配-----------------------------------
func repeatedSubstringPattern(s string) bool {
	sDouble := s + s
	L := len(sDouble)
	if strStr(sDouble[1:L-1], s) == -1 {
		return false
	}
	return true
}

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
