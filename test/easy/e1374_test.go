package test

import (
	"fmt"
	"math/rand"
	"testing"
)

const BASE string = "qwertyuioplkjhgfdsazxcvbnm"
const BASE_LEN int = 26

func generateTheString(n int) string {
	rec := make([]int, 26)
	for n > 0 {
		cur := rand.Intn(26)
		if rec[cur] != 0 {
			continue
		}
		dec := rand.Intn(n)
		for dec&1 == 0 {
			dec = rand.Intn(n)
			fmt.Println(dec)
		}
		rec[cur] = dec
		n -= dec
	}
	res := make([]byte, 26)
	i := 0
	for idx, cnt := range rec {
		for cnt > 0 {
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
