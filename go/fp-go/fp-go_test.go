package fp_go

import (
	"fmt"
	A "github.com/IBM/fp-go/array"
	N "github.com/IBM/fp-go/number"
	"testing"
)

func TestBasicFpGoFunctionForArray(t *testing.T) {
	input := []int{1, 2, 3, 4}
	red := A.Reduce(N.MonoidSum[int]().Concat, 0)(input)
	fmt.Println(red)
	fld := A.Fold(N.MonoidSum[int]())(input)
	fmt.Println(fld)
}

func TestBasicFunctionForMap(t *testing.T) {
	f := func(i int) int {
		return i * 2
	}

	input := []int{1, 2, 3, 4}
	res1 := make([]int, 0, len(input))
	for _, i := range input {
		res1 = append(res1, f(i))
	}
	fmt.Println(res1)
	res2 := A.Map(f)(input)
	fmt.Println(res2)
}
