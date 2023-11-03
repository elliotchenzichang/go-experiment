package _go

import (
	"fmt"
	"testing"
)

func TestStringConvertToInt(t *testing.T) {
	var a int = 1
	// can not use this string(a) to convert directly.
	str := string(rune(a))
	fmt.Println(str)
}
