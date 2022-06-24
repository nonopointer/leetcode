package test

import (
	"math"
	"testing"
)

func Test515(t *testing.T) {
	largestValues(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	arr := []*TreeNode{root}
	bfs(arr, &res)
	return res
}

func bfs(nodes []*TreeNode, res *[]int) {
	nexts := []*TreeNode{}
	currentMax := math.MinInt

	for _, node := range nodes {
		if node.Val > currentMax {
			currentMax = node.Val
		}
		if node.Left != nil {
			nexts = append(nexts, node.Left)
		}
		if node.Right != nil {
			nexts = append(nexts, node.Right)
		}
	}
	*res = append(*res, currentMax)
	if len(nexts) != 0 {
		bfs(nexts, res)
	}
}
