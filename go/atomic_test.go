package _go

import (
	"fmt"
	"reflect"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var value atomic.Value
	value.Store(1)
	value.CompareAndSwap(1, 2)
	v := value.Load()
	fmt.Println(value.Load())
	fmt.Println(reflect.TypeOf(v))
}
