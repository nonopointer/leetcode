package test

import (
	"fmt"
	"testing"
)

func Test1037(t *testing.T) {
	fmt.Println(isBoomerang([][]int{{73, 31}, {73, 19}, {73, 45}}))
}

func isBoomerang(points [][]int) bool {
	x1 := points[0][0] - points[1][0]
	x2 := (points[0][0] - points[2][0])
	y1 := points[0][1] - points[1][1]
	y2 := points[0][1] - points[2][1]
	return x1*y2 != x2*y1
}
