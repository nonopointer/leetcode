package test

import (
	"fmt"
	"sort"
	"testing"
)

func Test1200(t *testing.T) {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
}
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	min := 0x7fffffff

	var res [][]int
	for i := 0; i < len(arr)-1; i++ {
		var res1 []int
		tempMin := arr[i+1] - arr[i]
		if tempMin > min {
			continue
		}
		res1 = []int{arr[i], arr[i+1]}
		if tempMin == min {
			res = append(res, res1)
		}
		if tempMin < min {
			min = tempMin
			res = [][]int{res1}
		}
	}
	return res
}
