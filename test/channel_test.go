package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestChann(t *testing.T) {

	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	cnt := 1
	c1 <- struct{}{}
	go func() {

		for i := 0; i < 100; i++ {
			<-c1
			fmt.Println("go-1", cnt)
			cnt++
			c2 <- struct{}{}
		}
		wg.Done()
	}()

	go func() {

		for i := 0; i < 100; i++ {
			<-c2
			fmt.Println("go-2", cnt)
			cnt++
			c1 <- struct{}{}
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("end")
}
