package test

import (
	"fmt"
	"testing"
)

type WordFilter struct {
	Map map[string]int
}

func Constructor745(words []string) WordFilter {
	mp := map[string]int{}
	for i, word := range words {
		for j, n := 1, len(word); j <= n; j++ {
			for k := 1; k <= n; k++ {
				mp[word[:j]+" "+word[n-k:]] = i
			}
		}
	}
	return WordFilter{mp}
}

func (w *WordFilter) F(pref string, suff string) int {
	if v, ok := w.Map[pref+" "+suff]; ok {
		return v
	} else {
		return -1
	}
}

func Test745(t *testing.T) {
	wf := Constructor745([]string{"abc"})
	fmt.Println(wf.F("a","c"))
}
