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

type Solution710 struct {
	sign int
	base []bool
}

func Constructor710(n int, blacklist []int) Solution710 {
	base := make([]bool, n)
	for _, v := range blacklist {
		base[v] = true
	}
	res := Solution710{
		sign: n,
		base: base,
	}
	return res
}

func (s *Solution710) Pick() int {
	res := rand.Intn(s.sign)
	temp, cnt := 0, 0
	for cnt < res {
		if !s.base[temp] {
			cnt++
		}
		if cnt == res {
			return temp
		}
		temp++
		if temp == s.sign {
			temp = 0
		}
	}
	return temp
}
