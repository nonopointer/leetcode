package test

import (
	"fmt"
	"strings"
	"testing"
)

func Test1108(t *testing.T) {
	fmt.Println(defangIPaddr("1.1.1.1"))
	fmt.Println(defangIPaddr1("1.1.1.1"))
}

func defangIPaddr(address string) string {
	return strings.Join(strings.Split(address, "."), "[.]")
}

func defangIPaddr1(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
