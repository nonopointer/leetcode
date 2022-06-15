package test

import (
	"fmt"
	"sort"
	"testing"
)

func Test719(t *testing.T) {
	fmt.Println(smallestDistancePair([]int{1,3,5,6,9,9,4}, 3))
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
    return sort.Search(nums[len(nums)-1]-nums[0], func(mid int) bool {
        cnt := 0
        for j, num := range nums {
            i := sort.SearchInts(nums[:j], num-mid)
            cnt += j - i
        }
        return cnt >= k
    })
}
