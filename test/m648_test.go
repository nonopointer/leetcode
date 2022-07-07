package test

import (
	"fmt"
	"strings"
	"testing"
)

func replaceWords(dictionary []string, sentence string) string {
	dict := map[string]bool{}
	for _, str := range dictionary {
		dict[str] = true
	}
	ss := strings.Split(sentence, " ")
	for idx, word := range ss {
		for i := 1; i < len(word); i++ {
			if dict[word[:i]] {
				ss[idx] = word[:i]
				break
			}
		}
	}
	return strings.Join(ss, " ")
}

func Test648(t *testing.T) {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
}
