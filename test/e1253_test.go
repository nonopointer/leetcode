package test

import (
	"fmt"
	"testing"
)

func oddCells(m int, n int, indices [][]int) int {
	res := 0
	base := make([][]int, m)
	for i := 0; i < m; i++ {
		base[i] = make([]int, n)
	}

	for _, arr := range indices {
		for i := 0; i < n; i++ {
			base[arr[0]][i]++
		}
		for i := 0; i < m; i++ {
			base[i][arr[1]]++
		}
	}

	for _, a := range base {
		for _, b := range a {
			res += (b & 1)
		}
	}

	return res
}

func Test1253(t *testing.T) {
	fmt.Println(oddCells(2,3,[][]int{{0,1},{1,1}}))
}
