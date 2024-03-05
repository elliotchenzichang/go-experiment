package _go

import "testing"

func TestSelectRandom(t *testing.T) {
	var ch = make(chan int)
	go func() {
		for {
			ch <- 1
		}
	}()

	for {
		select {
		case <-ch:
			t.Logf("case 1")
		case <-ch:
			t.Logf("case 2")
		}
	}
}
