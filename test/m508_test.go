package test

import (
	"math"
	"testing"
)

func Test508(t *testing.T) {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var mm = make(map[int]int)
var max = math.MinInt

func findFrequentTreeSum(root *TreeNode) []int {

	dfs(root)
	res := make([]int, 0)
	// fmt.Println(mm)

	for k, v := range mm {
		if v == max {
			res = append(res, k)
		}
	}
	return res
}

func dfs(cur *TreeNode) int {
	cnt := cur.Val
	if cur.Left != nil {
		cnt += dfs(cur.Left)
	}
	if cur.Right != nil {
		cnt += dfs(cur.Right)
	}

	if v, ok := mm[cnt]; ok {
		mm[cnt] = v + 1
	} else {
		mm[cnt] = 1
	}
	if mm[cnt] > max {
		max = mm[cnt]
	}
	return cnt
}
