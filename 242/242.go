package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "anagram"
	t := "nagara"
	fmt.Println(isAnagram(s, t))
}

// map表方法
/*
func isAnagram(s string, t string) bool {
	//创建储存记录的map表
	sTable := make(map[int32]int, 26)
	tTable := make(map[int32]int, 26)

	//每次遍历到记录个数+1
	for _, val := range s {
		sTable[val]++
	}
	for _, val := range t {
		tTable[val]++
	}

	//如果字母个数都不一样，不是字母异位词
	if len(sTable) != len(tTable) {
		return false
	}

	//比较两个记录表是否相同
	for i, _ := range sTable {
		if sTable[i] != tTable[i] {
			return false
		}
	}
	return true

}
*/

// 排序方法
func isAnagram(s string, t string) bool {
	//StringToSlice
	S := []rune(s)
	T := []rune(t)

	//进行排序
	sort.Slice(S, func(i, j int) bool {
		return S[i] < S[j]
	})
	sort.Slice(T, func(i, j int) bool {
		return T[i] < T[j]
	})

	//比较排序后字符串是否相同
	if string(S) == string(T) {
		return true
	}
	return false
}
