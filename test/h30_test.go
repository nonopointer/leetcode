package test

import (
	"fmt"
	"testing"
	"time"
)

func Test30(t *testing.T) {
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
	fmt.Println(findSubstring1("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}

func findSubstring(s string, words []string) []int {
	start := time.Now().UnixNano()
	res := make([]int, 0)

	wordCnt := len(words)
	if wordCnt == 0 {
		return res
	}
	wordSize := len(words[0])
	totalSize := wordCnt * wordSize
	lex := len(s)
	l, r := 0, totalSize-1
	mw := make(map[string]int)

	for _, v := range words {
		mw[v]++
	}
	// fmt.Println(mw, r, lex)

	for r < lex {

		flag := true
		mm := make(map[string]int)

		for i := 0; i < wordCnt; i++ {
			str := s[l+wordSize*i : l+wordSize*(i+1)]
			if _, ok := mw[str]; !ok {
				break
			}
			mm[str]++
		}
		

		for k, v := range mw {
			if mm[k] != v {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, l)
		}
		l++
		r++
	}
	fmt.Println("cost:", time.Now().UnixNano()-start)

	return res
}

func findSubstring1(s string, words []string) []int {
	start := time.Now().UnixNano()
	res, basem := []int{}, map[string]int{}
	for _, str := range words {
		basem[str] += 1
	}
	wlen := len(words[0])
	for k := 0; k < wlen; k++ {
		m := map[string]int{}
		for _, str := range words {
			m[str] += 1
		}
		start := k
		for i := k; i+wlen-1 < len(s); i += wlen {
			str := s[i : i+wlen]
			if _, ok := basem[str]; !ok {
				for j := start; j < i; j += wlen {
					m[s[j:j+wlen]] += 1
				}
				start = i + wlen
				continue
			}
			num, ok := m[str]
			if !ok {
				j := start
				for ; j < i && s[j:j+wlen] != str; j += wlen {
					m[s[j:j+wlen]] += 1
				}
				start = j + wlen
			} else if num == 1 {
				delete(m, str)
			} else {
				m[str]--
			}
			if len(m) == 0 {
				res = append(res, start)
			}
		}
	}
	fmt.Println("cost:", time.Now().UnixNano()-start)
	return res
}
