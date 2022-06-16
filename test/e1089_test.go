package test

import (
	"fmt"
	"testing"
)

func Test1089(t *testing.T) {
	duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0})
}

func duplicateZeros(arr []int) {
	lex := len(arr)
	arr2 := make([]int, lex)
	idx := 0
	for i := 0; i < lex && idx < lex; i++ {
		arr2[idx] = arr[i]
		idx++
		if arr[i] == 0 && idx < lex {
			arr2[idx] = 0
			idx++
		}
	}
	// fmt.Println(arr)
	// arr = arr2

	for i, v := range arr2 {
		arr[i] = v
	}
	// fmt.Println(arr)
}
