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

type Slice[T int | float32 | float64] []T

func TestBasicGenericTypeDefinition(t *testing.T) {
	var a Slice[int] = []int{1, 2, 3}
	var b Slice[float32] = []float32{1.1, 2.2, 3.3}
	var c Slice[float64] = []float64{1.1, 2.2, 3.3}
	fmt.Println(a[0])
	fmt.Println(b[0])
	fmt.Println(c[0])
}

func TestDefineMapWithGenericType(t *testing.T) {
	type GenericMap[Key int | string, Value int | string] map[Key]Value
	var myMap GenericMap[int, string] = map[int]string{
		1: "Elliot",
		2: "Chen",
	}
	fmt.Println(myMap[1])
}

type StructureWithGenericType[T int64 | string] struct {
	Data T
}

func (f *StructureWithGenericType[T]) Print() {
	fmt.Println(f.Data)
}

func TestDefineStructureWithGenericType(t *testing.T) {
	var obj = &StructureWithGenericType[int64]{
		Data: 3,
	}
	obj.Print()
}
