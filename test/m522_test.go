package test

import (
	"fmt"
	"testing"
)

func Test522(t *testing.T) {
	fmt.Println(findLUSlength([]string{"aba", "cdc", "eae"}))
}

func findLUSlength(strs []string) int {
	n, ans := len(strs), -1
	for i := 0; i < n; i++ {
		if len(strs[i]) <= ans {
			continue
		}
		ok := true
		for j := 0; j < n && ok; j++ {
			if i == j {
				continue
			}
			if check(strs[i], strs[j]) {
				ok = false
			}
		}
		if ok {
			ans = len(strs[i])
		}
	}
	return ans
}

func check(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if m < n {
		return false
	}
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = f[i-1][j-1]
			}
			if f[i][j] < f[i-1][j] {
				f[i][j] = f[i-1][j]
			}
			if f[i][j] < f[i][j-1] {
				f[i][j] = f[i][j-1]
			}
			if f[i][j] == n {
				return true
			}
		}
	}
	return false
}
