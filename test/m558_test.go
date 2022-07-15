package test

import (
	"fmt"
	"testing"
)

type Node558 struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node558
	TopRight    *Node558
	BottomLeft  *Node558
	BottomRight *Node558
}

func intersect(t1 *Node558, t2 *Node558) *Node558 {

	if t1.IsLeaf && t2.IsLeaf {
		if t1.Val {
			return t1
		} else if t2.Val {
			return t2
		} else {
			return t1
		}
	}
	ans := &Node558{}
	var topLeftTemp1, topRightTemp1, bottomLeftTemp1, bottomRightTemp1 *Node558
	if t1.IsLeaf {
		topLeftTemp1, topRightTemp1, bottomLeftTemp1, bottomRightTemp1 = t1, t1, t1, t1
	} else {
		topLeftTemp1, topRightTemp1, bottomLeftTemp1, bottomRightTemp1 = t1.TopLeft, t1.TopRight, t1.BottomLeft, t1.BottomRight
	}
	var topLeftTemp2, topRightTemp2, bottomLeftTemp2, bottomRightTemp2 *Node558
	if t2.IsLeaf {
		topLeftTemp2, topRightTemp2, bottomLeftTemp2, bottomRightTemp2 = t2, t2, t2, t2
	} else {
		topLeftTemp2, topRightTemp2, bottomLeftTemp2, bottomRightTemp2 = t2.TopLeft, t2.TopRight, t2.BottomLeft, t2.BottomRight
	}
	ans.TopLeft = intersect(topLeftTemp1, topLeftTemp2)
	ans.TopRight = intersect(topRightTemp1, topRightTemp2)
	ans.BottomLeft = intersect(bottomLeftTemp1, bottomLeftTemp2)
	ans.BottomRight = intersect(bottomRightTemp1, bottomRightTemp2)
	a := ans.TopLeft.IsLeaf && ans.TopRight.IsLeaf && ans.BottomLeft.IsLeaf && ans.BottomRight.IsLeaf
	b := ans.TopLeft.Val && ans.TopRight.Val && ans.BottomLeft.Val && ans.BottomRight.Val
	c := ans.TopLeft.Val || ans.TopRight.Val || ans.BottomLeft.Val || ans.BottomRight.Val
	ans.IsLeaf = a && (b || !c)
	ans.Val = ans.TopLeft.Val
	if ans.IsLeaf {
		ans.TopLeft, ans.TopRight, ans.BottomLeft, ans.BottomRight = nil, nil, nil, nil
	}
	return ans

}

func Test558(t *testing.T) {
	fmt.Print(intersect(&Node558{}, &Node558{}))
}
