package test

import (
	"fmt"
	"testing"
)

type MovingAverage struct {
	base []int
	size int
	sum  int
}

/** Initialize your data structure here. */
func Constructor041(size int) MovingAverage {
	return MovingAverage{
		base: []int{},
		size: size,
		sum:  0,
	}
}

func (m *MovingAverage) Next(val int) float64 {
	lex := len(m.base)
	var head int
	if lex == m.size {
		head = m.base[0]
		m.base = m.base[1:]
		m.sum -= head
	}
	m.base = append(m.base, val)
	m.sum += val
	return float64(m.sum) / float64(len(m.base))
}

func TestFloat(t *testing.T) {
	// fmt.Println(float64(4) / float64(3))

	ma := Constructor041(3)
	fmt.Println(ma.Next(1))
	fmt.Println(ma.Next(10))
	fmt.Println(ma.Next(3))
	fmt.Println(ma.Next(5))

}
