package test

import "testing"

type RangeModule struct {
	mm map[int]int
}

func Constructor() RangeModule {
	return RangeModule{mm: make(map[int]int)}
}

func (r *RangeModule) AddRange(left int, right int) {
	for k, v := range r.mm {
		if k >= left && k <= right {
			left = v
		}
	}
}

func (r *RangeModule) QueryRange(left int, right int) bool {
	return false
}


func Test715(t *testing.T) {

}
