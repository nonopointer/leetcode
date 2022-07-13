package test

import (
	"fmt"
	"testing"
)

func Test735(t *testing.T) {
	fmt.Println(asteroidCollision([]int{-5, 10, -15}))
}

func asteroidCollision(asteroids []int) []int {
	lex := len(asteroids)
	res := make([]int, lex)
	res[0] = asteroids[0]
	idx := 0
	for i := 1; i < lex; i++ {
		fmt.Println("res : ", res[:idx+1])
		testsss(res, asteroids, &idx, i)
	}
	return res[:idx+1]
}

func testsss(res, asteroids []int, idx *int, i int) {
	if res[*idx] < 0 {
		*idx++
		res[*idx] = asteroids[i]
	} else {
		if confictTwoNum(res[*idx], asteroids[i]) {

			if res[*idx] > -asteroids[i] {
				return
			}
			if res[*idx] == -asteroids[i] {
				*idx--
				return
			}

			for *idx >= 0 && res[*idx] < -asteroids[i] {
				res[*idx] = asteroids[i]
				*idx--
				if *idx < 0 {
					*idx = 0
					break
				}
			}

		} else {
			*idx++
			res[*idx] = asteroids[i]
		}
	}
}

func confictTwoNum(a, b int) bool {
	return (a > 0) != (b > 0)
}
