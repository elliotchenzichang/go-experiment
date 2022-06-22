package nutsdb

import (
	"bytes"
	"fmt"
	"github.com/xujiajun/nutsdb"
	"math/rand"
	"testing"
	"time"
)

var nutsDB *nutsdb.DB

func init() {
	opts := nutsdb.DefaultOptions
	opts.Dir = "bench/nutsdb"
	opts.SyncEnable = false
	opts.EntryIdxMode = nutsdb.HintKeyAndRAMIdxMode
	var err error
	nutsDB, err = nutsdb.Open(opts)
	if err != nil {
		panic(err)
	}
}

func BenchmarkPutValue_NutsDB(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		nutsDB.Update(func(tx *nutsdb.Tx) error {
			err := tx.Put("test-bucket", GetKey(i), GetValue(), 0)
			if err != nil {
				panic(err)
			}
			return nil
		})
	}
}

func BenchmarkGetValue_NutsDB(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		nutsDB.View(func(tx *nutsdb.Tx) error {
			_, err := tx.Get("test-bucket", GetKey(i))
			if err != nil && err != nutsdb.ErrKeyNotFound {
				panic(err)
			}
			return nil
		})
	}
}

const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"

func init() {
	rand.Seed(time.Now().Unix())
}

func GetKey(n int) []byte {
	return []byte("test_key_" + fmt.Sprintf("%09d", n))
}

func GetValue() []byte {
	var str bytes.Buffer
	for i := 0; i < 512; i++ {
		str.WriteByte(alphabet[rand.Int()%36])
	}
	return []byte(str.String())
}
