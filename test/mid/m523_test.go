package test

import (
	"fmt"
	"testing"
)

func Test523(t *testing.T) {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 6}, 7))
}

func checkSubarraySum(nums []int, k int) bool {

	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	set := make(map[int]struct{})
	for i := 2; i <= n; i++ {
		set[sum[i-2]%k] = struct{}{}
		if _, ok := set[sum[i]%k]; ok {
			return true
		}
	}
	return false

	/**

	lex := len(nums)
	if lex < 2 {
		return false
	}

	sum := make([]int, lex+1)
	lexs := len(sum)
	for i := 1; i < lexs; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	mm := make(map[int]struct{})

	for i := 2; i < lexs; i++ {
		mm[sum[i-2]%k] = struct{}{}
		if _, ok := mm[sum[i]%k]; ok {
			return true
		}
	}

	return false
	*/
}
