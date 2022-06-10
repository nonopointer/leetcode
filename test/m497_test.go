package test

import (
	"fmt"
	"math/rand"
	"testing"
)

type Solution struct {

	rects [][]int
	cnts []int64
}

func Constructor(rects [][]int) Solution {

	// a := (rects[0][2] - rects[0][0]) * (rects[0][3] - rects[0][1])
	// b := (rects[1][2] - rects[1][0]) * (rects[1][3] - rects[1][1])
	// fmt.Println("a = ", a, " b =  ", b, "ap = ", float64(a)/float64(a+b))
	// maxS := 0
	// ps := make([]float64, 0)
	// es := make([]int, 0)
	// for _, v := range rects {
	// 	s := (v[2] - v[0]) * (v[3] - v[1])
	// 	maxS += s
	// 	es = append(es, s)
	// }
	// temp := maxS
	// for i := len(es) - 1; i >= 0; i-- {
	// 	ps = append(ps, float64(temp)/float64(maxS))
	// 	temp -= es[i]
	// }

	return Solution{
		rects: rects,
	}
}

func (s *Solution) Pick() []int {
	// x = rand.Intn(s.d)
	// y = rand.Intn(s.h)
	var x int
	var y int
	// p := rand.Float64()
	// fmt.Println("p = ", p, " , ap = ", s.ap)

	// for i, v := range s.ps {
	// if p <= v {
	points := s.rects[rand.Intn(len(s.rects))]
	x = rand.Intn(points[2]-points[0]+1) + points[0]
	y = rand.Intn(points[3]-points[1]+1) + points[1]
	// }
	// }

	// if p > s.ap {
	// 	// x = int(math.Ceil(rand.Float64()*float64(s.bx2-s.bx1))) + s.bx1
	// 	// y = int(math.Ceil(rand.Float64()*float64(s.by2-s.by1))) + s.by1

	// 	// x = rand.Intn(s.bx2-s.bx1+1) + s.bx1
	// 	// x = rand.Intn(s.by2-s.by1+1) + s.bx1
	// } else {
	// 	// x = int(math.Ceil(rand.Float64()*float64(s.ax2-s.ax1))) + s.ax1
	// 	// y = int(math.Ceil(rand.Float64()*float64(s.ay2-s.ay1))) + s.ay1
	// 	x = rand.Intn(s.ax2-s.ax1+1) + s.bx1
	// 	x = rand.Intn(s.ay2-s.ay1+1) + s.bx1
	// }

	return []int{x, y}
}

func Test497(t *testing.T) {
	// rand.Seed(time.Now().Unix())
	// fmt.Println(rand.Intn(10))
	s := Constructor([][]int{{1, 1, 5, 5}})
	fmt.Println(s.Pick())
}
