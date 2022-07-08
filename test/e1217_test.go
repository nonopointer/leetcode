package test

import "testing"

func Test1217(t *testing.T) {
	minCostToMoveChips([]int{2, 2, 2, 3, 3})
}

func minCostToMoveChips(position []int) int {
	cnt := 0
	lex := len(position)
	for i := 0; i < lex; i++ {
		cnt += (position[i] & 1)
	}
	cnt2 := lex - cnt
	if cnt2 < cnt {
		return cnt2
	}
	return cnt
}
