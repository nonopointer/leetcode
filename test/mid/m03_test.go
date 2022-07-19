package test

import (
	"fmt"
	"testing"
)

func TestMd5Test(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcar"))
}

func lengthOfLongestSubstring(s string) int {
	base := make([]int, 128)
	for i := 0; i < 128; i++ {
		base[i] = -1
	}
	max, l, len := 0, 0, len(s)
	for i := 0; i < len; i++ {
		c := s[i]
		if base[c]+1 > l {
			l = base[c] + 1
		}
		if i-l+1 > max {
			max = i - l + 1
		}
		base[c] = i
	}

	return max
}
