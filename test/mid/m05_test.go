package test

import (
	"fmt"
	"testing"
)

func Test05(t *testing.T) {
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	max := 0
	var str string
	flag := -1
	for i := range s {
		len, sub := checkOne(i, s)
		if max < len {
			max = len
			str = sub
			flag = 1
		}

		len, sub = checkTwo(i, s)
		if max < len {
			max = len
			str = sub
			flag = 0
		}
	}
	var rs []byte
	rs = make([]byte, 0)
	if flag == 0 {
		lex := len(str)
		for i := lex - 1; i >= 0; i-- {
			rs = append(rs, str[i])
		}
		for i := 0; i < lex; i++ {
			rs = append(rs, str[i])
		}
		return string(rs)
	} else {
		lex := len(str)
		for i := lex - 1; i > 0; i-- {
			rs = append(rs, str[i])
		}
		for i := 0; i < lex; i++ {
			rs = append(rs, str[i])
		}
		return string(rs)
	}

}

func checkOne(i int, s string) (int, string) {
	bs := make([]byte, 0)
	bs = append(bs, s[i])
	for l, r := i-1, i+1; l >= 0 && r < len(s); {
		if s[l] == s[r] {
			bs = append(bs, s[l])
			l--
			r++
		} else {
			break
		}
	}
	return len(bs)*2 - 1, string(bs)
}
func checkTwo(i int, s string) (int, string) {
	bs := make([]byte, 0)
	// bs = append(bs, s[i])
	for l, r := i, i+1; l >= 0 && r < len(s); {
		if s[l] == s[r] {
			bs = append(bs, s[l])
			l--
			r++
		} else {
			break
		}
	}
	return len(bs) * 2, string(bs)
}
