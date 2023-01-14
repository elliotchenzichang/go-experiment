package _go

import (
	"fmt"
	"math"
	"testing"
)

func TestNumber(t *testing.T) {
	var num = math.MaxUint32 + 2
	fmt.Println("num is ", num)
	fmt.Println("convert to uint32 is ", uint32(num))
}
