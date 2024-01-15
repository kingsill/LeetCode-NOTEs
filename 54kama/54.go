package main

import "fmt"

// acm模式

/*
func main() {
	//以byte切片的方式获取输入
	var str []byte

	//获取输入
	_, err := fmt.Scanln(&str)
	if err != nil {
		return
	}

	//将要插入的number先存为[]byte，方便使用
	number := []byte{'n', 'u', 'm', 'b', 'e', 'r'}

	var record []byte

	for _, i2 := range str {

		//这里比较需要使用''，实际对应的数值是不一样的
		if i2 >= '0' && i2 <= '9' {
			record = append(record, number...)
		} else {
			record = append(record, i2)
		}

	}

	//将[]byte转换为string，方便后续输出
	s := string(record)

	fmt.Println(s)
}
*/

// 使用原[]byte
func main() {
	//以byte切片的方式获取输入
	var str []byte

	//获取输入
	_, err := fmt.Scanln(&str)
	if err != nil {
		return
	}

	//将要插入的number先存为[]byte，方便使用
	number := []byte{'n', 'u', 'm', 'b', 'e', 'r'}

	for i := 0; i < len(str); i++ {

		//这里比较需要使用''，实际对应的数值是不一样的
		if str[i] >= '0' && str[i] <= '9' {
			str = append(str[:i], append(number, str[i+1:]...)...)
			i += 6 - 1
		}

	}

	//将[]byte转换为string，方便后续输出
	c := string(str)

	fmt.Println(c)
}
