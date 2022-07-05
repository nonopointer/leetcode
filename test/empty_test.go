package test

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestEm(t *testing.T) {

	x := []int{
		1,
		2,
	}
	_ = x

	var s1 struct{}
	var s2 struct{}
	fmt.Printf("size of s1 is %d\n", unsafe.Sizeof(s1))
	fmt.Printf("p1 = %p , p2 = %p , p1==p2: %v\n", &s1, &s2, &s1 == &s2)

	strChan := make(chan string, 3)
	signChan := make(chan struct{}, 1)  //接收数据信号
	signChan1 := make(chan struct{}, 1) //操作完成信号

	//receive
	go func() {
		<-signChan
		for v := range strChan {
			fmt.Println("received value is ", v)
		}
		signChan1 <- struct{}{}
	}()

	go func() {
		for _, v := range []string{"hello", "cl", "world"} {
			strChan <- v
		}
		close(strChan)
		signChan <- struct{}{}
	}()

	<-signChan1
	fmt.Println("Over !")
}

func TestNew(t *testing.T) {
	s := new([]int)
	fmt.Printf("%p\n", s)
	*s = append(*s, 1)
	fmt.Println(s, *s, len(*s), cap(*s))
	fmt.Printf("%p\n%p\n", s, &(*s)[0])

	var a [1]int
	a[0] = 1
	// fmt.Println(a)
	fmt.Printf("a\t\n%p\n%p\n", &a, &a[0])
	var a2 = [...]int{1, 2, 3}
	fmt.Printf("a2\t\n%p\n%p\n", &a2, &a2[0])

	// scm:= sync.Map{}
}
