package test

import (
	"fmt"
	"testing"
	"time"
)

func TestChann(t *testing.T) {
	//
	//c1 := make(chan struct{}, 1)
	//c2 := make(chan struct{}, 1)
	//c3 := make(chan struct{}, 1)

	const NUM = 7
	cs := make([]chan struct{}, NUM)
	for i := range cs {
		cs[i] = make(chan struct{}, 1)
	}
	exitCh := make(chan struct{}, 1)
	cnt := 1

	for i, j := 0, 0; i < NUM; i++ {
		go func(i int) {
			for {
				<-cs[i]
				if cnt > 111 {
					exitCh <- struct{}{}
					break
				}
				fmt.Println(i, cnt)
				cnt++
				if j == NUM-1 {
					j = 0
				} else {
					j++
				}
				cs[j] <- struct{}{}
			}
		}(i)
	}
	cs[0] <- struct{}{}
	fmt.Println("end")
	<-exitCh
	fmt.Println("Over")
	//select {
	//case <-exitCh:
	//	fmt.Println("Over")
	//}
}

func TestCh1(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 271828
	select {
	case ch <- 31415926:
		fmt.Println("通道的值为:", <-ch)
		fmt.Println("channel vaule is:", <-ch)
	default:
		fmt.Println("通道已经被阻塞,即已经满了")
	}
}

func TestCh2(t *testing.T) {
	i := 0
	ch := make(chan string, 0)
	defer close(ch)
	go func() {
	CuiStartLoop: //不加也可以,与后面break后的 CuiStartLoop相呼应,作为循环体的标识
		for {
			time.Sleep(time.Second * 1)
			fmt.Println(time.Now().Unix())
			fmt.Println("当前i的值为:", i)
			i++
			select {
			case m := <-ch:
				fmt.Println("输出为:", m)
				break CuiStartLoop
			default:
				fmt.Println("执行了default代码块")
			}
		}
	}()
	time.Sleep(time.Second * 4)
	ch <- "stop"
}
