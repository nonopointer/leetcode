package test

import (
	"fmt"
	"math/rand"
	"testing"
)

const BASE string = "qwertyuioplkjhgfdsazxcvbnm"
const BASE_LEN int = 26

func generateTheString(n int) string {
	rec := make([]int, BASE_LEN)
	res := make([]byte, n)
	n += 1
	for n > 1 {
		cur := rand.Intn(BASE_LEN)
		if rec[cur] != 0 {
			continue
		}
		dec := rand.Intn(n)
		for dec&1 == 0 {
			dec = rand.Intn(n)
			//fmt.Println(dec)
		}
		rec[cur] = dec
		n -= dec
	}
	i := 0
	for idx, cnt := range rec {
		for cnt > 0 {
			fmt.Println(i, idx)
			res[i] = BASE[idx]
			cnt--
			i++
		}
	}
	return string(res)
}

func Test1374(t *testing.T) {
	fmt.Println(generateTheString(6))
}

func TestRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Println(rand.Intn(10))
	}
}
