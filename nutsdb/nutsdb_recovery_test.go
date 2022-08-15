package nutsdb

import (
	"github.com/xujiajun/nutsdb"
	"testing"
)

var db *nutsdb.DB
var dbDir = "bench/nutsdb"
var opt nutsdb.Options

func init() {
	opt = nutsdb.DefaultOptions
	opt.Dir = dbDir
	opt.SyncEnable = false
	opt.EntryIdxMode = nutsdb.HintKeyAndRAMIdxMode
	opt.RWMode = nutsdb.FileIO
	var err error
	db, err = nutsdb.Open(opt)
	if err != nil {
		panic(err)
	}
}

// 准备测试需要的 500MB 数据
func prepareDbData() {
	defer db.Close()
	for i := 0; i < 2000000; i++ {
		err := db.Update(func(tx *nutsdb.Tx) error {
			err := tx.Put(GetBucket(), GetKey(i), GetValue(), 0)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return
		}
	}
}

func TestPrepareData(t *testing.T) {
	prepareDbData()
}

// go test -bench BenchmarkNutsDbMMapRecovery -run none -benchmem -cpuprofile cpuprofile_mmap.out -memprofile memprofile_mmap.out
func BenchmarkNutsDbMMapRecovery(b *testing.B) {
	opt.RWMode = nutsdb.MMap
	_, err := nutsdb.Open(opt)
	if err != nil {
		panic(err)
	}
}

// go test -bench BenchmarkNutsDbFileIORecovery -run none -benchmem -cpuprofile cpuprofile_fileio.out -memprofile memprofile_fileio.out
func BenchmarkNutsDbFileIORecovery(b *testing.B) {
	opt.RWMode = nutsdb.FileIO
	_, err := nutsdb.Open(opt)
	if err != nil {
		panic(err)
	}
}
