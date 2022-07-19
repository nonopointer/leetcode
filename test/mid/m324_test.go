package test

import (
	"sort"
	"testing"
)

func Test324(t *testing.T) {
	wiggleSort([]int{1, 3, 34, 455, 3})
}

/***
按题目条件我们想偶数坐标位填小一些的数，奇数坐标位填大一些的数，想到排序后分成两部分。
但是这两部分的分界处可能是一样大的，比如[4,5,5,6]的分界在5。
为了保证一样大的数会被错开，我们可以将前部分倒序填入，同时为了保证一定比它大，后部分也要倒序填入。
*/

func wiggleSort(nums []int) {
	n := len(nums)
	cp := append([]int{}, nums...)
	sort.Ints(cp)
	for idx, i, j := 0, (n+1)/2-1, n-1; idx < n; {
		nums[idx] = cp[i]
		idx++
		if idx < n {
			nums[idx] = cp[j]
			j--
			idx++
		}
		i--
	}
	// fmt.Println(nums)
}
