package test

import (
	"fmt"
	"testing"
)

func TestOff029(t *testing.T) {
	fmt.Println(insert(&Nodeoff029{},1))
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

type Nodeoff029 struct {
	Val  int
	Next *Nodeoff029
}

func insert(aNode *Nodeoff029, x int) *Nodeoff029 {
	if aNode == nil {
		t := &Nodeoff029{
			Val: x,
		}
		t.Next = t
		return t
	}
	if aNode.Next == aNode {
		aNode.Next = &Nodeoff029{
			Val:  x,
			Next: aNode,
		}
		return aNode
	}
	var temp *Nodeoff029
	temp = aNode.Next

	min, max := aNode.Val, aNode.Val
	for temp != aNode {
		if max < temp.Val {
			max = temp.Val
		}
		if min > temp.Val {
			min = temp.Val
		}
		temp = temp.Next
	}

	if x <= min || x >= max {
		for {
			if temp.Val == max {
				temp.Next = &Nodeoff029{
					Val:  x,
					Next: temp.Next,
				}
				break
			}
			temp = temp.Next
		}
	} else {
		for {
			if temp.Val <= x && temp.Next.Val >= x {
				temp.Next = &Nodeoff029{
					Val:  x,
					Next: temp.Next,
				}
				break
			}
			temp = temp.Next
		}
	}
	return aNode
}
