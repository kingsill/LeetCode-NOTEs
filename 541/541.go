package main

import "fmt"

func main() {
	s := "12312312312"
	k := 2
	c := reverseStr(s, k)
	fmt.Println(c)
}

func reverseStr(s string, k int) string {
	L := len(s)

	str := []byte(s)

	record := []byte{}

	n := L / k

	for i := 0; i < n; i++ {
		temp := str[k*i : (i+1)*k]

		//假如是奇数，则需要进行倒序
		if i%2 == 0 {
			reverseString(temp)
		}

		record = append(record, temp...)
	}

	temp := str[n*k:]
	if n%2 == 0 {
		reverseString(temp)
	}

	record = append(record, temp...)
	return string(record)
}

func reverseString(s []byte) {
	L := len(s)
	//双指针
	left, right := 0, L-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
