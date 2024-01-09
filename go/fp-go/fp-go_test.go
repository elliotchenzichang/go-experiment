package fp_go

import (
	"fmt"
	A "github.com/IBM/fp-go/array"
	N "github.com/IBM/fp-go/number"
	"testing"
)

func TestBasicFpGoFunction(t *testing.T) {
	input := []int{1, 2, 3, 4}
	red := A.Reduce(N.MonoidSum[int]().Concat, 0)(input)
	fmt.Println(red)
	fld := A.Fold(N.MonoidSum[int]())(input)
	fmt.Println(fld)
}
