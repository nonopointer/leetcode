package test

import (
	"fmt"
	"strconv"
	"testing"
)

func Test241(t *testing.T) {
	fmt.Println(diffWaysToCompute("1-1+2"))
}

func diffWaysToCompute(expression string) []int {
	if v, ok := isNum(expression); ok {
		return []int{v}
	}
	res := []int{}
	for idx, c := range expression {
		tempC := string(c)
		if tempC == "+" || tempC == "-" || tempC == "*" {
			left := diffWaysToCompute(expression[:idx])
			right := diffWaysToCompute(expression[idx+1:])

			for _, leftNum := range left {
				for _, rightNum := range right {
					var result int
					switch tempC {
					case "+":
						result = leftNum + rightNum
					case "-":
						result = leftNum - rightNum
					case "*":
						result = leftNum * rightNum
					}
					res = append(res, result)
				}
			}

		}
	}
	return res
}

func isNum(s string) (int, bool) {
	if v, err := strconv.Atoi(s); err == nil {
		return v, true
	}
	return 0, false
}
