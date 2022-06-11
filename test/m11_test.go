package test

import (
	"fmt"
	"testing"
)

func Test11(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	lex := len(height) - 1
	l, r := 0, lex
	var rs int

	if height[l] < height[r] {
		rs = lex * height[l]
	} else {
		rs = lex * height[r]
	}
	temp := 0
	for l < r {
		lex--
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		if height[l] < height[r] {
			temp = lex * height[l]
		} else {
			temp = lex * height[r]
		}
		if temp > rs {
			rs = temp
		}
	}

	return rs
}
