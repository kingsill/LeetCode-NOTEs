package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))

}

func evalRPN(tokens []string) int {

	L := len(tokens)

	for {
	term:
		tokens2 := transfer(tokens)
		for i, v := range tokens {

			if v == "+" || v == "-" || v == "*" || v == "/" {
				switch v {
				case "+":
					result := tokens2[i-2] + tokens2[i-1]
					tokens2, tokens = CustAppend(result, i, L, tokens2, tokens)
					goto term
				case "-":
					result := tokens2[i-2] - tokens2[i-1]
					tokens2, tokens = CustAppend(result, i, L, tokens2, tokens)
					goto term
				case "*":
					result := tokens2[i-2] * tokens2[i-1]
					tokens2, tokens = CustAppend(result, i, L, tokens2, tokens)
					goto term
				case "/":
					result := tokens2[i-2] / tokens2[i-1]
					tokens2, tokens = CustAppend(result, i, L, tokens2, tokens)
					goto term
				}

			}

		}
		if len(tokens2) == 1 {
			return tokens2[0]
		}

	}
}

func transfer(tokens []string) (tokens2 []int) {
	tokens2 = make([]int, 0)
	for _, v := range tokens {
		num, _ := strconv.Atoi(v)
		tokens2 = append(tokens2, num)
	}
	return tokens2

}

func CustAppend(result int, i int, L int, tokens2 []int, tokens []string) ([]int, []string) {
	s := strconv.Itoa(result)
	if i == L-1 {
		tokens2 = append(tokens2[:i-2], result)

		tokens = append(tokens[:i-2], s)
	} else {
		tokens2 = append(append(tokens2[:i-2], result), tokens2[i+1:]...)
		tokens = append(append(tokens[:i-2], s), tokens[i+1:]...)
	}
	return tokens2, tokens
}
