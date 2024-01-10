package _go

import (
	"fmt"
	"github.com/bytedance/sonic"
	"strings"
	"testing"
)

func TestSonic(t *testing.T) {
	var o = map[string]interface{}{}
	var r = strings.NewReader(`{"a":"b"}{"1":"2"}`)
	var dec = sonic.ConfigDefault.NewDecoder(r)
	dec.Decode(&o)
	dec.Decode(&o)
	fmt.Printf("%+v", o)
}

func TestSonic_MarshalType(t *testing.T) {
	type Stu struct {
		Name string
		Age  int
	}

	var stuStr = `{"name":"Elliot", "age": 18}`
	var stu *Stu
	err := sonic.Unmarshal([]byte(stuStr), &stu)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", stu)

}
