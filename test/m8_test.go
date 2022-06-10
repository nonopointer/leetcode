package test

import (
	"fmt"
	"strconv"
	"testing"
)

func myAtoi(s string) int {
	fmt.Println(s)
	res := make([]byte, 0)
	flag := -1 // -1 not start 0 start 1 end
	var zf int64 = 1
	for i, v := range s {
		fmt.Printf("%d\t%v\n", i, res)
		if flag == 1 {
			break
		}

		if flag == -1 && (v < '0' || v > '9') && v != '+' && v != '-' && v != ' ' {
			flag = 1
			break
		}

		if flag == 0 && (v < '0' || v > '9') {
			flag = 1
			break
		}

		if v >= '0' && v <= '9' {
			flag = 0
			if len(res) == 0 && v == '0' {
				continue
			}
			// fmt.Println("len res is ", len(res))
			res = append(res, byte(v))
		}

		if v == '+' {
			flag = 0
		}
		if v == '-' {
			flag = 0
			zf = -1
		}
		//|| v == '-'
	}

	rs := string(res)
	n, _ := strconv.ParseInt(rs, 0, 64)
	n = n * zf
	// fmt.Println(n)
	if n < -(1 << 31) {
		return -(1 << 31)
	}

	if n > 0x7fffffff {
		return 0x7fffffff
	}
	return int(n)
}

func Test8(t *testing.T) {
	fmt.Println(myAtoi("  0000000000012345678"))

	// fmt.Println(-(1<<31))
	// fmt.Println(0x7fffffff)
	// fmt.Println(1<<32)
}
