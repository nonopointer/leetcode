package test

import (
	"fmt"
	"testing"
)

func Test890(t *testing.T) {
	fmt.Println(findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
}
func findAndReplacePattern(words []string, pattern string) []string {

	rs := make([]string, 0)
	for _, v := range words {
		m1 := make(map[byte]byte, len(pattern))
		m2 := make(map[byte]byte, len(pattern))
		flag := true
		for i, c := range v {
			if m1[byte(c)] == 0 {
				m1[byte(c)] = pattern[i]
			} else if m1[byte(c)] != pattern[i] {
				flag = false
				break
			}
			if m2[pattern[i]] == 0 {
				m2[pattern[i]] = byte(c)
			} else if m2[pattern[i]] != byte(c) {
				flag = false
				break
			}
		}
		if flag {
			rs = append(rs, v)
		}
	}
	return rs
}
