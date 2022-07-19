package test

import (
	"fmt"
	"testing"
)

type MagicDictionary map[string]bool

func Constructor676() MagicDictionary {
	return MagicDictionary{}
}

func (m *MagicDictionary) BuildDict(dictionary []string) {
	for _, v := range dictionary {
		(*m)[v] = true
	}
}

const base = "qwertyuioplkjhgfdsazxcvbnm"

func (m *MagicDictionary) Search(searchWord string) bool {
	// temp :=
	for i := range searchWord {
		if checkword(i, searchWord, m) {
			return true
		}
	}
	return false
}

func checkword(idx int, word string, dic *MagicDictionary) bool {
	bs := []byte(word)
	old := bs[idx]
	for i := range base {
		if old == base[i] {
			continue
		}
		bs[idx] = base[i]
		if (*dic)[string(bs)] {
			return true
		}
	}
	return false
}

func Test676(t *testing.T) {
	m := Constructor676()
	m["aa"] = true
	fmt.Println(m)
	fmt.Println(m.Search("ab"))
	fmt.Println(m.Search("aa"))
}
