package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var n = 0
var m int32 = 0

func TestThread(t *testing.T) {

	sign := make(chan struct{}, 2)

	go func() {
		for i := 0; i < 1000000; i++ {
			n++
		}
		sign <- struct{}{}
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			n++
		}
		sign <- struct{}{}
	}()
	<-sign
	<-sign
	fmt.Println(n)
}

func TestT2(t *testing.T) {
	const WGCNT = 5
	wg := sync.WaitGroup{}
	wg.Add(WGCNT)

	for i := 0; i < WGCNT; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000000; i++ {
				atomic.AddInt32(&m, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(m)
	time.Sleep(1 * time.Second)
	fmt.Println(m)
}

func TestArray(t *testing.T) {
	arr := [3]int{1, 2, 3}
	array(&arr)
	fmt.Println(arr)
	fmt.Println("***********")
	array2(arr)
	fmt.Println(arr)
}

func array(arr *[3]int) {
	(*arr)[0] = 9
	fmt.Println(*arr)
}

func array2(arr [3]int) {
	arr[0] = 8
	fmt.Println(arr)
}

func TestA(t *testing.T) {
	var arr = [5]int{1, 2}
	arr[2] = 3
	arr[3] = 4
	fmt.Println(arr) // 输出：[1 4 3 4 0]
}
