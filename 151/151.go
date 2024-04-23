package main

import "fmt"

func main() {

	origin := "a good   example"
	s := reverseWords(origin)
	fmt.Println(s)
}

/*func reverseWords(s string) string {
	WordList := []string{}
	left, right := 0, 0
	L := len(s)
	//fmt.Println(L)

	for i, i2 := range s {
		//去除重复的空格
		if i > 0 && s[i-1] == ' ' && i2 == ' ' {
			left++
			continue
		}

		//不为空格时右指针移动，并且如果最后一位也为字母时添加到存储表
		if i2 != ' ' {
			right = i
			if i == L-1 {
				WordList = append(WordList, s[left:right+1])
			}
		} else { //当为空格时，将当前的左右指针对应的单词放入存储表，同时给左指针赋值
			if i != 0 {
				WordList = append(WordList, s[left:right+1])
			}
			left = i + 1
		}
	}
	S := WordList[0]
	for i, s2 := range WordList {
		if i != 0 {
			S = s2 + " " + S
		}
	}
	return S
}*/
/*
func reverseWords(s string) string {
	WordList := []string{}
	left := 0
	L := len(s)
	//fmt.Println(L)

	for i, i2 := range s {
		//去除重复的空格
		if i > 0 && s[i-1] == ' ' && i2 == ' ' {
			left++
			continue
		}

		//不为空格时右指针移动，并且如果最后一位也为字母时添加到存储表
		if i2 != ' ' {
			if i == L-1 {
				WordList = append(WordList, s[left:i+1])
			}
		} else { //当为空格时，将当前的左右指针对应的单词放入存储表，同时给左指针赋值
			if i != 0 {
				WordList = append(WordList, s[left:i])
			}
			left = i + 1
		}
	}
	S := WordList[0]
	for i, s2 := range WordList {
		if i != 0 {
			S = s2 + " " + S
		}
	}
	return S
}*/

// 重来版本
func reverseWords(s string) string {
	str := []byte(s)
	str = cutHead(str)
	reverse(str)
	str = cutHead(str)
	str = removeSpace(str)
	L := len(str)
	right := 0
	left := 0
	for ; right < L; right++ {
		if str[right]-' ' == 0 {
			reverse(str[left:right])
			left = right + 1
		}
	}

	reverse(str[left:L])

	return string(str)
}

func cutHead(str []byte) []byte {
	for i, i2 := range str {
		if i2-' ' != 0 {
			return str[i:]
		}
	}
	return str
}

func reverse(str []byte) {
	L := len(str)
	left := 0
	right := L - 1
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
}

func removeSpace(str []byte) []byte {
	index := 0
	for index < len(str) {
		if str[index]-' ' == 0 && str[index-1]-' ' == 0 {
			str = append(str[:index], str[index+1:]...)
			index--
		}
		index++
	}
	return str
}
