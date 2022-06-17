package test

import "testing"

func TestOff029(t *testing.T) {

}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		t := &Node{
			Val: x,
		}
		t.Next = t
		return t
	}
	if aNode.Next == aNode {
		aNode.Next = &Node{
			Val:  x,
			Next: aNode,
		}
		return aNode
	}
	var temp *Node
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
				temp.Next = &Node{
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
				temp.Next = &Node{
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
