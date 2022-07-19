package test

import (
	"fmt"
	"testing"
)

func TestOffer091(t *testing.T) {
	fmt.Println(minCost([][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}))
}

func minCost(costs [][]int) int {
	for i := 1; i < len(costs); i++ {
		costs[i][0] += Min(costs[i-1][1], costs[i-1][2])
		costs[i][1] += Min(costs[i-1][0], costs[i-1][2])
		costs[i][2] += Min(costs[i-1][0], costs[i-1][1])
	}
	return Min(costs[len(costs)-1][0], Min(costs[len(costs)-1][1], costs[len(costs)-1][2]))

}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
