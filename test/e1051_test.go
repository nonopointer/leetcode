package test

import (
	"fmt"
	"testing"
)

func Test1051(t *testing.T) {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}

func heightChecker(heights []int) int {
	lex := len(heights)
	target := make([]int, lex)
	copy(target, heights)

	var temp int

	for i := range target {
		for j := i + 1; j < lex; j++ {
			if target[i] > target[j] {
				temp = target[j]
				target[j] = target[i]
				target[i] = temp
			}
		}
	}
	res := 0

	for i := range target {
		if target[i] != heights[i] {
			res++
		}
	}
	return res
}
