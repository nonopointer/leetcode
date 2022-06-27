package test

import (
	"math/rand"
	"testing"
)

func Test710(t *testing.T) {
	/*

	 XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

	*/
}

type Solution struct {
	sign int
	base []bool
}

func Constructor(n int, blacklist []int) Solution {
	base := make([]bool, n)
	for _, v := range blacklist {
		base[v] = true
	}
	res := Solution{
		sign: n,
		base: base,
	}
	return res
}

func (this *Solution) Pick() int {
	res := rand.Intn(this.sign)
	temp, cnt := 0, 0
	for cnt < res {
		if !this.base[temp] {
			cnt++
		}
		if cnt == res {
			return temp
		}
		temp++
		if temp == this.sign {
			temp = 0
		}
	}
	return temp
}
