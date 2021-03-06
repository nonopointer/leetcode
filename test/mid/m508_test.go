package test

import (
	"fmt"
	"math"
	"testing"
)

func Test508(t *testing.T) {
	root := &TreeNode508{
		Val: 1,
	}
	fmt.Println(findFrequentTreeSum(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode508 struct {
	Val   int
	Left  *TreeNode508
	Right *TreeNode508
}

func findFrequentTreeSum(root *TreeNode508) []int {

	var mm = make(map[int]int)
	var max = math.MinInt

	dfs(root, mm, &max)
	res := make([]int, 0)

	for k, v := range mm {
		if v == max {
			res = append(res, k)
		}
	}
	return res
}

func dfs(cur *TreeNode508, mm map[int]int, max *int) int {
	cnt := cur.Val
	if cur.Left != nil {
		cnt += dfs(cur.Left, mm, max)
	}
	if cur.Right != nil {
		cnt += dfs(cur.Right, mm, max)
	}

	if v, ok := mm[cnt]; ok {
		mm[cnt] = v + 1
	} else {
		mm[cnt] = 1
	}
	if mm[cnt] > *max {
		*max = mm[cnt]
	}
	return cnt
}
