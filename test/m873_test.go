package test

import (
	"fmt"
	"testing"
)

func Test873(t *testing.T) {
	// fmt.Println(lenLongestFibSubseq1([]int{1, 3, 7, 11, 12, 14, 18}))
	fmt.Println(lenLongestFibSubseq([]int{1, 3, 7, 11, 12, 14, 18}))
}

func lenLongestFibSubseq(arr []int) (ans int) {
    n := len(arr)
    indices := make(map[int]int, n)
    for i, x := range arr {
        indices[x] = i
    }
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for i, x := range arr {
        for j := n - 1; j >= 0 && arr[j]*2 > x; j-- {
            if k, ok := indices[x-arr[j]]; ok {
                dp[j][i] = max(dp[k][j]+1, 3)
                ans = max(ans, dp[j][i])
            }
        }
    }
    return
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}


func lenLongestFibSubseq1(arr []int) int {
	lex := len(arr)
	last := arr[lex-1]
	mm := map[int]int{}
	for a, b := 1, 1; b <= last; {
		mm[b] = 1
		temp := a + b
		a = b
		b = temp
	}

	fmt.Println(mm)
	res := 0
	for _, v := range arr {
		res += mm[v]
	}
	return res
}
