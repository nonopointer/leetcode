package test

import (
	"fmt"
	"math"
	"testing"
)

var N, INF = 55, math.MinInt
var f = make([][][]int, 2*N)

func init() {
	for i := 0; i < 2*N; i++ {
		f[i] = make([][]int, N)
		for j := 0; j < N; j++ {
			f[i][j] = make([]int, N)
		}
	}
}
func cherryPickup(grid [][]int) int {
	// int n = grid.length;
	n := len(grid)
	for k := 0; k <= 2*n; k++ {
		for i1 := 0; i1 <= n; i1++ {
			for i2 := 0; i2 <= n; i2++ {
				f[k][i1][i2] = INF
			}
		}
	}
	f[2][1][1] = grid[0][0]
	for k := 3; k <= 2*n; k++ {
		for i1 := 1; i1 <= n; i1++ {
			for i2 := 1; i2 <= n; i2++ {
				j1, j2 := k-i1, k-i2
				if j1 <= 0 || j1 > n || j2 <= 0 || j2 > n {
					continue
				}
				A, B := grid[i1-1][j1-1], grid[i2-1][j2-1]
				if A == -1 || B == -1 {
					continue
				}
				a, b, c, d := f[k-1][i1-1][i2], f[k-1][i1-1][i2-1], f[k-1][i1][i2-1], f[k-1][i1][i2]
				//  t := Math.max(Math.max(a, b), Math.max(c, d)) + A;
				t := Max(Max(a, b), Max(c, d)) + A
				if i1 != i2 {
					t += B
				}
				f[k][i1][i2] = t
			}
		}
	}
	return Max(f[2*n][n][n], 0)
}
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Test741(t *testing.T) {
	fmt.Println(cherryPickup([][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}))
}
