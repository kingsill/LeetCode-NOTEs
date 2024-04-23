package main

import "fmt"

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	//1.判断长度，只有偶数可能满足

	L := len(s)
	if L%2 != 0 {
		return false
	}

	//使用切片作为栈
	List := make([]byte, L)

	//define the map of open bracket
	Dictionary := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	//栈的指针
	index := 0

	for i := 0; i < L; i++ {

		//if open bracket,add correspondind close bracket to List
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			List[index] = Dictionary[s[i]]
			index++
		} else {

			//close bracket more
			if index-1 < 0 {
				return false
			}

			//dont corresponding
			if s[i] != List[index-1] {
				return false
			} else {
				index--
			}
		}

		//open bracket more
		if i == L-1 && index != 0 {
			return false
		}

	}
	return true

}
