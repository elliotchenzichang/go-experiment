package gen_struct_in_runtime

import (
	"fmt"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	pe := NewBuilder().
		AddString("Name").
		AddInt64("Age").
		Build()
	p := pe.New()
	p.SetString("Name", "你好")
	p.SetInt64("Age", 32)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%T，%+v\n", p.Interface(), p.Interface())
	fmt.Printf("%T，%+v\n", p.Addr(), p.Addr())
}
