package main

import (
	"fmt"
)

func main() {
	s := "aacabdkacaa"
	r := longestPalindrome(s)
	fmt.Println(r)
}

func longestPalindrome(s string) string {
	length := len(s) //字符串长度
	var X, Y, M = 0, 0, 0
	table := make([][]int, length) //用于保存信息的表格
	for i := 0; i < length; i++ {
		table[i] = make([]int, length)
	}
	//fmt.Println(table)

	////首先将对角线的元素都标记为回文字符串
	//for i := 0; i < length; i++ {
	//	table[i][i] = 0
	//}
	//不需要以上定义，初始化表格的过程中自动赋值为0
	//为了方便，我们直接通过初始化将所有标记为0，意味都为回文子串，通过挨个修正来得到正确结果
	//其目的是为了省略对没有左下角数据的字符的额外判断

	//从第一行开始遍历
	for x := 0; x < length; x++ {
		for y := x + 1; y < length; y++ {
			u := y - x - 1 //大X为实际对应的表格内坐标
			//状态转移方程中的对于子子串的判断在没有左下角的数据时是不启用的
			if s[u] == s[y] && table[u+1][y-1] == 0 { //如果是回文子串，记录其长度
				if y-u >= M { //判断当前回文子串是否为最长回文子串
					M = y - u
					X = u
					Y = y
				}
				//table[u][y] = 0
			} else {
				table[u][y] = 1
			}

		}
	}
	//for _, ints := range table {
	//	fmt.Println(ints)
	//}
	n := s[X : Y+1]
	return n
}
