package main

func main() {

}

func reverseStr(s string, k int) string {
	L := len(s)

	str := []byte(s)

	record := [][]byte{}

	n := L / k

	for i := 0; i < n; i++ {
		record = append(record, str[:])
	}

	reverseString(str)

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
