package test

import (
	"fmt"
	"sort"
	"testing"
)

func Test15(t *testing.T) {
	for _, v := range threeSum([]int{-1, 0, 1, 2, -1, -4}) {
		fmt.Println(v)
	}
}

func threeSum(nums []int) [][]int {
	lists := make([][]int, 0)
	//排序
	sort.Ints(nums)
	//双指针
	lex := len(nums)
	for i := 0; i < lex; i++ {
		if nums[i] > 0 {
			return lists
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		curr := nums[i]
		L, R := i+1, lex-1
		for L < R {
			tmp := curr + nums[L] + nums[R]
			if tmp == 0 {
				list := make([]int, 0)
				list = append(list, curr)
				list = append(list, nums[L])
				list = append(list, nums[R])
				lists = append(lists, list)
				for L < R && nums[L+1] == nums[L] {
					L++
				}
				for L < R && nums[R-1] == nums[R] {
					R--
				}
				L++
				R--
			} else if tmp < 0 {
				L++
			} else {
				R--
			}
		}
	}
	return lists
}
