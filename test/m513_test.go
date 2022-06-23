package test

import (
	"fmt"
	"testing"
)

func Test513(t *testing.T) {
	fmt.Println(findBottomLeftValue(&TreeNode513{Val: 1}))
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
	s1 := [6]int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:]
	// fmt.Println(len(s2))
	fmt.Printf("s1=%d\ns2=%p\ns3=%p\ns4=%p\n", s1, &s2, &s1[0], &s1[1])

	fmt.Println("---------------------")

	c1 := []int{1, 2, 3, 4}
	fmt.Printf("%p\n%p\n%p\n", &c1, &c1[0], &c1[1])
	fmt.Println("---------------------")

	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Printf("a=%p\ta=%v\nb=%p\tb=%v\n", &a, a, &b, b)
	a = append(a[:2], a[3:]...)
	fmt.Printf("a=%p\ta=%v\nb=%p\tb=%v\n", &a, a, &b, b)

	fmt.Println("-------copy--------------")

	x := []int{1, 2, 3, 4, 5}
	y := make([]int, len(x))
	copy(y, x)
	fmt.Printf("x=%p\tx=%v\ny=%p\ty=%v\n", &x, x, &y, y)
	x = append(x[:2], x[3:]...)
	fmt.Printf("x=%p\tx=%v\ny=%p\ty=%v\n", &x, x, &y, y)

}
