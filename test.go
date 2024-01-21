package main

import "fmt"

func reverse(strByte []byte, l, r int) {
	for l < r {
		strByte[l], strByte[r] = strByte[r], strByte[l]
		l++
		r--
	}
}

func main() {
	var str string
	var target int

	fmt.Scanln(&target)
	fmt.Scanln(&str)
	strByte := []byte(str)

	reverse(strByte, 0, len(strByte)-1)
	reverse(strByte, 0, target-1)
	reverse(strByte, target, len(strByte)-1)

	fmt.Printf(string(strByte))
}
