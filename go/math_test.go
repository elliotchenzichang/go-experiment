package _go

import (
	"fmt"
	"math"
	"testing"
)

func TestPow10(t *testing.T) {
	for i := 1; i < 53; i++ {
		fmt.Println(math.Pow10(i))
	}
}
