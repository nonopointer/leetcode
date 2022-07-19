package test

import (
	"fmt"
	"testing"
)

// type MyCalendar struct {
// 	link Zone
// }

// type Zone struct {
// 	zone [2]int
// 	next *Zone
// }

// func Constructor() MyCalendar {
// 	return MyCalendar{}
// }

// func (this *MyCalendar) Book(start int, end int) bool {
// 	node := [2]int{start, end}
// 	if (this.link == Zone{}) {
// 		this.link = Zone{zone: node}
// 		return true
// 	}
// 	temp := &this.link
// 	var pre *Zone
// 	for (temp != &Zone{}) {
// 		z := temp.zone
// 		fmt.Println(z,start,end)
// 		if z[0] <= end && z[0] >= start {
// 			if z[1] > end {
// 				node[1] = z[1]
// 			}
// 			// temp.next = temp.next.next
// 			pre.next = temp.next// delete
// 			if z[0] < end {
// 				return false
// 			}
// 		}
// 		if start > z[0] {
// 			newNode := &Zone{
// 				zone: node,
// 				next: temp.next,
// 			}
// 			temp.next = newNode
// 		}
// 		if start > z[1] {
// 			break
// 		}
// 		pre = temp
// 		temp = temp.next
// 	}
// 	return true
// }

type MyCalendar struct {
	zones [][]int
}

func Constructor729() MyCalendar {
	return MyCalendar{zones: [][]int{}}
}

func (m *MyCalendar) Book(start int, end int) bool {
	zone := []int{start, end}
	for _, z := range m.zones {
		if z[0] < end && z[1] > start {
			return false
		}
	}
	m.zones = append(m.zones, zone)
	return true
}
func Test729(t *testing.T) {
	c := Constructor729()
	fmt.Println(c.Book(1, 2))
	fmt.Println(c.Book(2, 4))
	fmt.Println(c.Book(5, 6))
	fmt.Println(c.Book(3, 4))
}
