package main

import "fmt"

func main() {
	str := "hello"
	s := []byte(str)
	reverseString(s)
	fmt.Println(string(s))
}

/*func reverseString(s []byte) {
	L := len(s)
	//双指针
	left, right := 0, L-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
*/

// 重来
func reverseString(s []byte) {
	L := len(s)

	left := 0
	right := L - 1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

}
