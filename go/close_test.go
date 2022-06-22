package _go

import (
	"fmt"
	"testing"
	"time"
)

func TestClose(t *testing.T) {
	ch := make(chan int, 5)

	go func() {
		for {
			select {
			case i := <-ch:
				fmt.Println(i)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch) // 关闭ch
	time.After(3 * time.Second)
}
