package main

func main() {
	ransomNote := "ab"
	magzine := "aa"
	canConstruct(ransomNote, magzine)
}

// 使用map进行记录
func canConstruct(ransomNote string, magazine string) bool {
	record := make(map[int32]int)

	for _, i2 := range magazine {

		record[i2]++
	}

	for _, i2 := range ransomNote {
		if record[i2] <= 0 {
			return false
		}
		record[i2]--

	}
	return true

}
