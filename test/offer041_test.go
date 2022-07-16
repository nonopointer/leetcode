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
func Constructor(size int) MovingAverage {
	return MovingAverage{
		base: []int{},
		size: size,
		sum:  0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	lex := len(this.base)
	var head int
	if lex == this.size {
		head = this.base[0]
		this.base = this.base[1:]
		this.sum -= head
	}
	this.base = append(this.base, val)
	this.sum += val
	return float64(this.sum) / float64(len(this.base))
}

func TestFloat(t *testing.T) {
	// fmt.Println(float64(4) / float64(3))

	ma := Constructor(3)
	fmt.Println(ma.Next(1))
	fmt.Println(ma.Next(10))
	fmt.Println(ma.Next(3))
	fmt.Println(ma.Next(5))

}
