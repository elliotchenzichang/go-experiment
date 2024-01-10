package _go

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect_Base(t *testing.T) {
	var num = 1
	fmt.Println(reflect.TypeOf(num))
	v := reflect.ValueOf(&num)
	v.Elem().SetInt(100)
	fmt.Println(num)
}
