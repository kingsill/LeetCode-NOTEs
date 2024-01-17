package main

import "fmt"

func main() {

	origin := "  hello world  "
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
}
