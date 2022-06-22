package test

import (
	"fmt"
	"testing"
)

func Test513(t *testing.T) {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode513 struct {
	Val   int
	Left  *TreeNode513
	Right *TreeNode513
}

func findBottomLeftValue(root *TreeNode513) int {
	queue := []*TreeNode513{root}
	var res int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		res = node.Val
	}
	return res
}

func TestSilce(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:]
	fmt.Println(len(s2))
}
