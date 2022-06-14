package test

import (
	"fmt"
	"testing"
)

func Test498(t *testing.T) {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	//[1 2 4 7 5 3 3 3 3
	//[1,2,4,7,5,3,6,8,9]

}

func findDiagonalOrder(mat [][]int) []int {
	var m, n int
	if m = len(mat); m > 0 {
		n = len(mat[0])
	}
	if m == 0 || n == 0 {
		return []int{}
	}

	res := make([]int, m*n)
	row, col := 0, 0
	lexRes := len(res)
	for idx := 0; idx < lexRes; idx++ {
		res[idx] = mat[row][col]

		if (row&1)^(col&1) == 0 {
			if col == n-1 {
				row++
			} else if row == 0 {
				col++
			} else {
				row--
				col++
			}
		} else {
			if row == m-1 {
				col++
			} else if col == 0 {
				row++
			} else {
				row++
				col--
			}
		}
	}
	return res
}
