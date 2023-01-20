package _go

import (
	"fmt"
	"testing"
)

type Numeric interface {
	int | float32 | float64
}

func TestGeneric(t *testing.T) {
	type Node[T Numeric] struct {
		value T
	}

	n := Node[int]{value: 5}
	fmt.Println(n.value)

}

func max[T Numeric](a, b T) bool {
	// ...
	return a > b
}

func TestGeneric_Method(t *testing.T) {
	max(1, 2)
}
