package main

import (
	"fmt"
)

func main() {
	n := 4
	fmt.Println(generateMatrix(n))
}

// 初版，我们从中心点开始
/*func generateMatrix(n int) [][]int {
	//1.nXn矩阵
	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
		//fmt.Println(table[i]) //确认生成n*n
	}

	//2.确定中心点位置
	//n=1,	(0,0)
	//n=2,	(1,0)
	//n=3,	(1,1)
	//n=4,	(2,1)
	var x, y = 0, 0 //中心点坐标
	x = n / 2
	y = (n - 1) / 2
	fmt.Println("n:", n, "	x:", x, "	y:", y)
	table[x][y] = n * n

	//3.填充

	var char = 0  //都是从横行开始移动, 横竖移动标志位
	var char2 int //到底是向左还是向右
	var char3 int //到底是向上还是向下

	if n%2 == 1 {
		char3 = 1
		char2 = -1
	} else {
		char3 = -1
		char2 = 1
	}
	var c = 1
	for t, i := 1, n*n-1; i >= 0; { //总共填入n*n个数 ,t为当前运行的次数，i为本次填入的数据

		for e := 0; e < 2; e++ {

			tem := (c + 1) / 2 //1,1,2,2,3,3

			if char == 0 { //横着动
				for tem > 0 {
					y = y + char2
					if y == n+1 || y == -1 {
						return table
					}
					table[x][y] = i
					i--
					t++

					tem--
				}
				char2 = -char2
				char = 1
			} else { //竖着动
				for tem > 0 {
					x = x + char3
					if x == n+1 || x == -1 {
						return table
					}
					table[x][y] = i
					i--
					t++
					if t == n*n {
						return table
					}
					tem--
				}
				char3 = -char3
				char = 0

			}
			c++
		}
	}

	return table
}
*/
func generateMatrix(n int) [][]int {

	//1.nXn矩阵
	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
		//fmt.Println(table[i]) //确认生成n*n
	}

	i := 1
	x, y := 0, 0
	c := n - 1

	for ; c >= 0; c = c - 2 {
		//for z := 0; z < 4; z++ {
		table[x][y] = i
		for b := 0; b < c; b++ {
			table[x][y] = i
			if i == n*n {
				return table
			}
			y++
			i++

		}
		for b := 0; b < c; b++ {
			table[x][y] = i

			x++
			i++
		}
		for b := 0; b < c; b++ {
			table[x][y] = i
			y--
			i++
		}
		for b := 0; b < c-1; b++ {
			table[x][y] = i
			x--
			i++
		}
		y++
		//}
	}
	return table
}
