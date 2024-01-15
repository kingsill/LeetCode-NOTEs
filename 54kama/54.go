package main

import "fmt"

// acm模式
func main() {
	var str []byte
	_, err := fmt.Scanln(&str)
	if err != nil {
		return
	}

	numTable := map[byte]int{49: 1, 50: 2, 52: 3, 53: 4, 54: 5, 55: 6, 56: 7, 57: 8, 58: 9, 48: 0}

	var result string

	for _, i2 := range str {
		//fmt.Println(i2)
		if _, ok := numTable[i2]; ok {
			result = result + "number"
		} else {
			result = result + string(i2)
		}

	}
	fmt.Println(result)
}
