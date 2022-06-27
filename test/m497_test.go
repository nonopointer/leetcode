package test

import (
	"fmt"
	"math/rand"
	"testing"
)

type Solution497 struct {
	rects [][]int
}

func Constructor497(rects [][]int) Solution497 {

	return Solution497{
		rects: rects,
	}
}

func (s *Solution497) Pick() []int {
	var x int
	var y int
	points := s.rects[rand.Intn(len(s.rects))]
	x = rand.Intn(points[2]-points[0]+1) + points[0]
	y = rand.Intn(points[3]-points[1]+1) + points[1]

	return []int{x, y}
}

func Test497(t *testing.T) {
	s := Constructor497([][]int{{1, 1, 5, 5}})
	fmt.Println(s.Pick())
}
