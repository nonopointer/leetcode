package test

import (
	"fmt"
	"testing"
	"time"
)

func Test730(t *testing.T) {
	fmt.Println(countPalindromicSubsequences("abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba"))
}

func countPalindromicSubsequences2(S string) int {
	start := time.Now().UnixNano()
	defer fmt.Println("cost time : ", time.Now().UnixNano()-start, "start : ", start)
	const mod = 1e9 + 7
	//O(1)时间查到当前位置前后的第一个abcd
	pre := make([][]int, len(S))
	last := []int{-1, -1, -1, -1}
	for i, char := range S {
		offset := char - 'a'
		last[offset] = i
		pre[i] = make([]int, 4)
		copy(pre[i], last)
	}

	next := make([][]int, len(S))
	last = []int{len(S), len(S), len(S), len(S)}
	for i := len(S) - 1; i >= 0; i-- {
		offset := S[i] - 'a'
		last[offset] = i
		next[i] = make([]int, 4)
		copy(next[i], last)
	}

	dp := make([][]int, len(S))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(S))
	}

	for i := len(S) - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < len(S); j++ {
			if S[i] != S[j] {
				//只取i + 只取j - ij都不取
				dp[i][j] = (dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1] + mod) % mod
			} else {
				dp[i][j] = (2 * dp[i+1][j-1]) % mod
				l, r := next[i+1][S[i]-'a'], pre[j-1][S[i]-'a']
				if l < r {
					//2个S[i]，[l+1,r-1]里的回文串已经被aa包裹住了，去掉重复
					dp[i][j] = (dp[i][j] - dp[l+1][r-1] + mod) % mod
				} else if l == r {
					//1个S[i]，多aa
					dp[i][j] += 1
				} else {
					//0个S[i]，多a和aa
					dp[i][j] += 2
				}
			}
		}
	}
	return dp[0][len(S)-1] % mod
}

func countPalindromicSubsequences(s string) int {
	//非常巧妙的DP题
	//dp[i][j]表示S[i...j]的不同回文子序列个数
	//怎么去掉重复部分?
	//如果S[i] != S[j], dp[i][j] = dp[i + 1][j] + dp[i][j - 1] - dp[i + 1][j - 1] (中间计算两遍的减掉)
	//如果S[i] == S[j], 找在i右边的S[i]相同的第一个位置left,在j右边的和S[j]相同的第一个位置right
	//通过left和right的关系,判断S[i + 1] ~ S[j - 1]有多少个S[i]
	//如果有0个,那么dp[i][j] = (dp[i + 1][j - 1] * 2 + 2)  其中的2表示单个S[i]以及S[i]+S[i]这两个情况
	//如果有1个,那么dp[i][j] = (dp[i + 1][j - 1] * 2 + 1)  其中的1表示S[i]+S[i]
	//如果有多个,那么dp[i][j] = (dp[i + 1][j - 1] * 2 - dp[left + 1][right - 1])
	start := time.Now().UnixNano()
	defer func() {
		end := time.Now().UnixNano()
		fmt.Println("cost time : ", end-start, "start : ", start, "end :", end)
	}()
	const mod = 1000000007
	n := len(s)
	//vector<vector<int>> dp(n + 1 , vector<int> (n + 1 , 0)) ;
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][i] = 1
	}
	for len := 2; len <= n; len++ {
		for i := 1; i <= n; i++ {
			j := i + len - 1
			if j > n {
				break
			}
			if s[i-1] == s[j-1] {
				left := i + 1
				right := j - 1
				for left < j && s[left-1] != s[i-1] {
					left++
				}
				for right > i && s[right-1] != s[i-1] {
					right--
				}
				if right < left {
					dp[i][j] = (dp[i+1][j-1]*2 + 2) % mod
				} else if right == left {
					dp[i][j] = (dp[i+1][j-1]*2 + 1) % mod
				} else {
					dp[i][j] = (dp[i+1][j-1]*2%mod + (mod - dp[left+1][right-1])) % mod
				}
			} else {
				dp[i][j] = (dp[i+1][j] + dp[i][j-1]) % mod
				dp[i][j] = (dp[i][j] + (mod-dp[i+1][j-1])%mod) % mod
			}
		}
	}
	return dp[1][n]
}
