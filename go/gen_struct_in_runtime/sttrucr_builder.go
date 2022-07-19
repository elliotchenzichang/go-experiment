package gen_struct_in_runtime

import (
	"errors"
	"reflect"
)

type Builder struct {
	fileId []reflect.StructField
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) AddField(field string, typ reflect.Type) *Builder {
	f := reflect.StructField{Name: field, Type: typ}
	b.fileId = append(b.fileId, f)
	return b
}

func (b *Builder) AddString(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(""))
}

func (b *Builder) AddBool(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(true))
}

func (b *Builder) AddInt64(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(int64(0)))
}

func (b *Builder) AddFloat64(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(float64(1.2)))
}

func (b *Builder) Build() *Struct {
	stu := reflect.StructOf(b.fileId)
	index := make(map[string]int)
	for i := 0; i < stu.NumField(); i++ {
		index[stu.Field(i).Name] = i
	}
	return &Struct{stu, index}
}

var (
	FieldNoExist = errors.New("field no exist")
)

func (in Instance) Field(name string) (reflect.Value, error) {
	if i, ok := in.index[name]; ok {
		return in.instance.Field(i), nil
	} else {
		return reflect.Value{}, FieldNoExist
	}
}

type Struct struct {
	typ reflect.Type
	// <fieldName : 索引> // 用于通过字段名称，从Builder的[]reflect.StructField中获取reflect.StructField
	index map[string]int
}

func (s Struct) New() *Instance {
	return &Instance{reflect.New(s.typ).Elem(), s.index}
}

type Instance struct {
	instance reflect.Value
	// <fieldName : 索引>
	index map[string]int
}

func (in *Instance) SetString(name, value string) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetString(value)
	}
}

func (in *Instance) SetBool(name string, value bool) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetBool(value)
	}
}

func (in *Instance) SetInt64(name string, value int64) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetInt(value)
	}
}

func (in *Instance) SetFloat64(name string, value float64) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetFloat(value)
	}
}
func (i *Instance) Interface() interface{} {
	return i.instance.Interface()
}

func (i *Instance) Addr() interface{} {
	return i.instance.Addr().Interface()
}
