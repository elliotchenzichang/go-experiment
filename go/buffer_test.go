package _go

import (
	"sync"
	"testing"
)

func BenchmarkPoolTest(b *testing.B) {
	pool := &sync.Pool{
		New: func() interface{} {
			buf := make([]byte, 500)
			return buf
		},
	}
	b.ResetTimer()
	b.SetParallelism(100)
	b.ReportAllocs()
	var bytes []byte
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bytes = pool.Get().([]byte)
			_ = len(bytes)
			pool.Put(bytes)
		}
	})
}

func BenchmarkByteTest(b *testing.B) {
	b.ResetTimer()
	b.SetParallelism(100)
	b.ReportAllocs()
	var bytes []byte
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bytes = make([]byte, 500)
			_ = len(bytes)
		}
	})
}
