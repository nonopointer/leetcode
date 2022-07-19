package test

import (
	"fmt"
	"sort"
	"testing"
)

type MyCalendarTwo struct {
	M map[int]int
	sortKey []int
}

func Constructor731() MyCalendarTwo {
	return MyCalendarTwo{M: map[int]int{}}
}

func (m *MyCalendarTwo) Book(start int, end int) bool {
	cnt := 0
	if _,ok := m.M[start];!ok {
		m.sortKey = append(m.sortKey, start)
	}
	if _,ok := m.M[end];!ok {
		m.sortKey = append(m.sortKey, end)
	}
	m.M[start]++
	m.M[end]--
	sort.Ints(m.sortKey)
	for _, v := range m.sortKey {
		cnt += m.M[v]
		if cnt > 2 {
			m.M[start]--
			m.M[end]++
			return false
		}
	}
	return true
}

func Test731(t *testing.T) {
	m := Constructor731()
	fmt.Println(m.Book(10, 20))
	fmt.Println(m.Book(50, 60))
	fmt.Println(m.Book(10, 40))
	fmt.Println(m.Book(5, 15))
	fmt.Println(m.Book(5, 10))
	fmt.Println(m.Book(25, 55))
}
