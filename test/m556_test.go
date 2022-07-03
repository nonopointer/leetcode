package test

import (
	"fmt"
	"testing"
)

func Test556(t *testing.T) {
	fmt.Println(nextGreaterElement(2147483486))
}

func nextGreaterElement(n int) int {
	baseArr := []int{}
	// temp := n
	for n > 0 {
		baseArr = append(baseArr, n%10)
		n = n / 10
	}

	for i := 1; i < len(baseArr); i++ {
		if baseArr[i] < baseArr[i-1] {
			return choose(baseArr, i)
		}
	}
	return -1
}

func choose(arr []int, idx int) int {
	temp := 10
	idxx := -1
	for i := 0; i < idx; i++ {
		if arr[i] < temp && arr[i] > arr[idx] {
			temp = arr[i]
			idxx = i
		}
	}
	swap := arr[idxx]
	arr[idxx] = arr[idx]
	arr[idx] = swap
	// fmt.Println(arr)

	for i := 0; i < idx; i++ {
		// fmt.Println("arr:", arr,"i:",i)
		for j := i + 1; j < idx; j++ {
			fmt.Println("arr:", arr)
			if arr[i] < arr[j] {
				// fmt.Println("arr:", arr)
				swap = arr[i]
				arr[i] = arr[j]
				arr[j] = swap
			}
		}
	}
	// fmt.Println(arr)
	res := 0
	for i := len(arr) - 1; i >= 0; i-- {
		res = res*10 + arr[i]
	}
	if res > 0x7fffffff {
		return -1
	}
	return res
}
