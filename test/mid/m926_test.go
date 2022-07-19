package test

import (
	"fmt"
	"testing"
)

func Test926(t *testing.T) {
	fmt.Println(minFlipsMonoIncr("00011001"))
	fmt.Println("-------------")
	fmt.Println(minFlipsMonoIncr("000110010"))
	fmt.Println("-------------")
	fmt.Println(minFlipsMonoIncr("0001100100"))
}

func minFlipsMonoIncr(s string) int {
	const zero = '0'
	baseZero := 0
	lex := len(s)
	cnt := 0
	rs := cnt

	for i := 0; i < lex; i++ {
		if s[i] == zero {
			cnt--
			baseZero++
			if rs > cnt {
				rs = cnt
			}
		} else {
			cnt++
		}
	}
	return rs + baseZero
}
