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
