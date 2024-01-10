package go_learn

import (
	"github.com/dolthub/swiss"
	"testing"
)

func TestSwiss_MapBase(t *testing.T) {
	m := swiss.NewMap[string, int](42)

	m.Put("foo", 1)
	m.Put("bar", 2)

	m.Iter(func(k string, v int) (stop bool) {
		println("iter", k, v)
		return false // continue
	})

	if x, ok := m.Get("foo"); ok {
		println(x)
	}
	if m.Has("bar") {
		x, _ := m.Get("bar")
		println(x)
	}

	m.Put("foo", -1)
	m.Delete("bar")

	if x, ok := m.Get("foo"); ok {
		println(x)
	}
	if m.Has("bar") {
		x, _ := m.Get("bar")
		println(x)
	}
}
