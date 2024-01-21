package main

import "fmt"

func main() {

	//1.使用递归？

	//2.从头开始遍历，但是从第k个开始写入

	/*	for i, b := range middle {

		}*/

	//3.直接修改middle切片
	//前K个单独切出来，再将其插入到后一部分的后面
	/*
			var k int
			var s string

			fmt.Scanln(&k)
			fmt.Scanln(&s)

			middle := []byte(s)

			//获取切片长度
			L := len(s)

		front := make([]byte, L-k)
			back := make([]byte, k)

			copy(front, middle[:L-k])
			copy(back, middle[L-k:])

			bytes := append(back, front...)

			result := string(bytes)

			fmt.Println(result)*/

	//4.先将全部字符倒叙，再在两部分分别进行倒序

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
func reverse(strByte []byte, l, r int) {
	for l < r {
		strByte[l], strByte[r] = strByte[r], strByte[l]
		l++
		r--
	}
}
