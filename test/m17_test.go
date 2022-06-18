package test

import (
	"fmt"
	"testing"
)

func Test17(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	mm := make(map[rune][]rune)

	mm['2'] = []rune{'a', 'b', 'c'}
	mm['3'] = []rune{'d', 'e', 'f'}
	mm['4'] = []rune{'g', 'h', 'i'}
	mm['5'] = []rune{'j', 'k', 'l'}
	mm['6'] = []rune{'m', 'n', 'o'}
	mm['7'] = []rune{'p', 'q', 'r', 's'}
	mm['8'] = []rune{'t', 'u', 'v'}
	mm['9'] = []rune{'w', 'x', 'y', 'z'}

	res := make([]string, 0)
	lex := len(digits)
	cs := make([][]rune, lex)
	for i, v := range digits {
		cs[i] = mm[v]
	}
	resMap := make(map[string]struct{})
	for _, vs := range cs {
		word := make([]rune, lex)
		idx := 0
		for _, v := range vs {
			fmt.Println(idx)
			word[idx] = v
			idx++
		}
		resMap[string(word)] = struct{}{}
	}
	for v := range resMap {
		res = append(res, v)
	}
	return res
}
