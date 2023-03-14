package _go

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var value atomic.Value
	value.Store(1)
	value.CompareAndSwap(1, 2)
	fmt.Println(value.Load())
}
