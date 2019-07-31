package main

import (
	"fmt"
	"strings"
	"time"
)

func searchMatrix(matrix [][]int, target int) bool {
	result := false
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		fmt.Println(1)
		return result
	}
	y := len(matrix)
	x := len(matrix[0])
	var limit int
	var isX bool
	if y > x {
		isX = false
		limit = x
	} else if y < x {
		isX = true
		limit = y
	} else if y == x {
		isX = true
		limit = x
	}

	lockLine := -1

	for i := 0; i < limit; i++ {
		if matrix[i][i] == target {
			fmt.Println(2)
			result = true
			break
			return result
		}
		if matrix[i][i] > target {
			fmt.Println(3)
			lockLine = i
			break
		}
	}

	fmt.Printf("x %d y %d %d lockLine %v isX\n", x, y, lockLine, isX)
	if lockLine == -1 {
		if isX {
			if y == 1 {
				for _, tt := range matrix[0] {
					if tt == target {
						fmt.Println("3.1")
						result = true
						break
					}
				}
			} else {
				for _, ty := range matrix {
					if ty[x-1] == target {
						fmt.Println(4)
						result = true
						break
					}
				}
			}
		} else {
			if x == 1 {
				for _, yy := range matrix {
					if yy[0] == target {
						fmt.Println("5.1")
						result = true
						break
					}
				}
			} else {
				for _, tx := range matrix[y-1] {
					if tx == target {
						fmt.Println(5)
						result = true
						break
					}
				}
			}
		}
	} else {
		// y
		for a := 0; a < lockLine; a++ {
			if matrix[a][lockLine] == target {
				fmt.Println(6)
				result = true
				break
			}
		}
		// z
		for b := 0; b < lockLine; b++ {
			if matrix[lockLine][b] == target {
				fmt.Println(7)
				result = true
				break
			}
		}
	}

	fmt.Println(8)
	return result
}

func main() {
	// test := [][]int{
	// 	[]int{1, 4, 7, 11, 15},
	// 	[]int{2, 5, 8, 12, 19},
	// 	[]int{3, 6, 9, 16, 22},
	// 	[]int{10, 13, 14, 17, 24},
	// 	[]int{18, 21, 23, 26, 30},
	// }

	// test := [][]int{
	// 	[]int{1,   2,   3,   4,   5},
	// 	[]int{6,   7,   8,   9,   10},
	// 	[]int{11, 12, 13, 14, 15},
	// 	[]int{16, 17, 18, 19, 20},
	// 	[]int{21, 22, 23, 24, 25},
	// }

	// fmt.Println(searchMatrix(test, 15))
	m, _ := time.ParseDuration("-100m")
	now := time.Now().Add(m)

	time.Sleep(5 * time.Second)

	diff := now.Sub(time.Now())
	fmt.Println(strings.Split(fmt.Sprintf("%v", diff), ".")[0] + "s")
}
