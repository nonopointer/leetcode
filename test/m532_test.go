package test

import (
	"fmt"
	"testing"
)

func Test532(t *testing.T) {
	fmt.Println(findPairs([]int{6, 3, 5, 7, 2, 3, 3, 8, 2, 4}, 2))
}

func findPairs(nums []int, k int) int {
	if k == 0 {
		return check0(nums)
	}
	mm := make(map[int]struct{})
	cnt := 0

	for _, num := range nums {
		if _, ok := mm[num]; !ok {
			mm[num] = struct{}{}
		}
	}

	for v := range mm {
		if _, ok := mm[v+k]; ok {
			cnt++
		}
	}
	return cnt
}

func check0(nums []int) int {
	mm := make(map[int]int)
	cnt := 0
	for _, num := range nums {
		if v, ok := mm[num]; ok {
			if v == 1 {
				cnt++
				mm[num] = 2
			}
		} else {
			mm[num] = 1
		}
	}
	return cnt
}
