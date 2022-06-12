package test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

type Solution11 struct {
	radius   float64
	x_center float64
	y_center float64
}

func Constructor11(radius float64, x_center float64, y_center float64) Solution11 {
	return Solution11{
		radius:   radius,
		x_center: x_center,
		y_center: y_center,
	}
}

func (s *Solution11) RandPoint() []float64 {
	r := math.Sqrt(rand.Float64()) * s.radius
	sin, cos := math.Sincos(rand.Float64() * 2 * math.Pi)
	return []float64{s.x_center + r*sin, s.y_center + r*cos}
}

func Test478(t *testing.T) {
	// s := Constructor(23.2, 354.2, 212.3)
	// fmt.Printf("%v\n", s.RandPoint())

	for i:=0 ;i<20;i++ {
		fmt.Println(math.Sqrt(rand.Float64()))
	}
}
