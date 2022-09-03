package nutsdb

import (
	"github.com/xujiajun/nutsdb"
	"math/rand"
	"testing"
)

var db *nutsdb.DB

func init() {
	opts := nutsdb.DefaultOptions
	opts.RWMode = nutsdb.FileIO
	opts.SyncEnable = true
	opts.Dir = "data"
	opts.EntryIdxMode = nutsdb.HintKeyAndRAMIdxMode
	db, _ = nutsdb.Open(opts)

	for i := 0; i < 100; i++ {
		db.Update(func(tx *nutsdb.Tx) error {
			tx.Put("test_bucket", GetKey(i), GetValue(), 0)
			return nil
		})
	}
}

func BenchmarkConcurrentRead(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			db.View(func(tx *nutsdb.Tx) error {
				for i := 0; i < 100; i++ {
					tx.Get("test_bucket", GetKey(i))
				}
				return nil
			})
		}
	})
}

func BenchmarkRead(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		db.View(func(tx *nutsdb.Tx) error {
			tx.Get("test_bucket", GetKey(rand.Intn(100)))
			return nil
		})
	}
}
