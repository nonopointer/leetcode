package test

import (
	"fmt"
	"testing"
)

type MyCalendarThree struct {
	max  int
	list [][]int
	cnts []int
}

func Constructor1() MyCalendarThree {
	return MyCalendarThree{max: 0, list: make([][]int, 0), cnts: make([]int, 0)}
}

func (m *MyCalendarThree) Book(start int, end int) int {
	cmax := 1
	for i, v := range m.list {
		// if (v[0] >= start && v[0] < end) ||
		// 	(v[0] <= start && v[1] > end) ||
		// 	(v[1] > start && v[1] < end) {
		if v[0] >= start && v[0] < end {
			m.cnts[i]++
			fmt.Println("cnts[", i, "]", "=", m.cnts[i], "max=", m.max)
			if len(m.cnts) > 2 {
				fmt.Println("------cnts[2] = ", m.cnts[2])
			}
			if m.max < m.cnts[i] {
				m.max = m.cnts[i]
			}
		}

		// if (v[0] <= start && v[0] > end) ||
		// 	(v[0] >= start && v[1] < end) ||
		// 	(v[1] < start && v[1] > end) {
		if v[0] <= start && v[1] > start {
			cmax++
		}

	}
	fmt.Println("cmax***", m.max, "***", cmax)
	if m.max < cmax {
		m.max = cmax
	}
	m.list = append(m.list, []int{start, end})
	m.cnts = append(m.cnts, cmax)
	return m.max
}

func TestXxx(t *testing.T) {
	c := Constructor1()
	//["MyCalendarThree","book","book","book","book","book","book"]
	//[[],[10,20],[50,60],[10,40],[5,15],[5,10],[25,55]]

	fmt.Println(c.Book(10, 20))
	fmt.Println(c.Book(50, 60))
	fmt.Println(c.Book(10, 40))
	fmt.Println(c.Book(5, 15))
	fmt.Println(c.Book(5, 10))
	fmt.Println(c.Book(25, 55))
}
