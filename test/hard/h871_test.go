package test

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

func Test871(t *testing.T) {
	fmt.Println(minRefuelStops(100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}))
}

type BiggestQueue struct {
	sort.IntSlice
}

func (b *BiggestQueue) Less(i, j int) bool {
	return b.IntSlice[i] > b.IntSlice[j]
}

func (b *BiggestQueue) Push(v interface{}) {
	b.IntSlice = append(b.IntSlice, v.(int))
}

func (b *BiggestQueue) Pop() interface{} {
	temp := b.IntSlice
	v := temp[len(temp)-1]
	b.IntSlice = temp[:len(temp)-1]
	return v
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	fuel, prev, bq, res := startFuel, 0, BiggestQueue{}, 0

	for i, n := 0, len(stations); i <= n; i++ {
		curr := target
		if i < n {
			curr = stations[i][0]
		}
		fuel -= curr - prev

		for fuel < 0 && bq.Len() > 0 {
			fuel += heap.Pop(&bq).(int)
			res++
		}
		if fuel < 0 {
			return -1
		}
		if i < n {
			heap.Push(&bq, stations[i][1])
			prev = curr
		}

	}

	return res
}
