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
