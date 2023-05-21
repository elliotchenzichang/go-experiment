package nutsdb

import (
	"fmt"
	"github.com/nutsdb/nutsdb"
	"log"
	"testing"
)

func TestNutsdb(t *testing.T) {
	opt := nutsdb.DefaultOptions
	opt.Dir = "tmp/nutsdb"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	key := []byte("key001")
	value := []byte("value001")
	bucket01 := "bucket001"
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put(bucket01, key, value, 0); err != nil {
				return err
			}
			return nil
		},
	); err != nil {
		log.Fatal(err)
	}
	defer func(db *nutsdb.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
}

func TestScan(t *testing.T) {
	opt := nutsdb.DefaultOptions
	opt.Dir = "tmp/nutsdb"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	var bucket = "bucket001"
	var data = map[string]string{
		"key001": "value001",
		"key002": "value002",
		"key003": "value003",
	}
	db.Update(func(tx *nutsdb.Tx) error {
		for key, value := range data {
			tx.Put(bucket, []byte(key), []byte(value), nutsdb.Persistent)
		}
		return nil
	})

	var prefixScanFunc = func(tx *nutsdb.Tx) error {
		fmt.Println("==================start prefix scan=======================")
		entries, _, _ := tx.PrefixScan(bucket, []byte("key"), 0, 1)
		for _, entry := range entries {
			fmt.Println(string(entry.Key), string(entry.Value))
		}
		fmt.Println("==================end prefix scan=======================")
		return nil
	}

	var prefixSearchScanFunx = func(tx *nutsdb.Tx) error {
		fmt.Println("==================start prefix search scan=======================")
		entries, _, _ := tx.PrefixSearchScan(bucket, []byte("key"), "00", 0, 2)
		for _, entry := range entries {
			fmt.Println(string(entry.Key), string(entry.Value))
		}
		fmt.Println("===================end prefix search scan=======================")
		return nil
	}

	var rangeScanFunc = func(tx *nutsdb.Tx) error {
		fmt.Println("==================start range scan=======================")
		entries, _ := tx.RangeScan(bucket, []byte("key001"), []byte("key002"))
		for _, entry := range entries {
			fmt.Println(string(entry.Key), string(entry.Value))
		}
		fmt.Println("===================end range scan=======================")
		return nil
	}

	var getAllFunc = func(tx *nutsdb.Tx) error {
		fmt.Println("==================start get all scan=======================")
		entries, _ := tx.GetAll(bucket)
		for _, entry := range entries {
			fmt.Println(string(entry.Key), string(entry.Value))
		}
		fmt.Println("==================end get all scan=======================")
		return nil
	}

	db.View(prefixScanFunc)
	db.View(prefixSearchScanFunx)
	db.View(rangeScanFunc)
	db.View(getAllFunc)
}

func TestViewData(t *testing.T) {
	db, err := nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir("/tmp/nutsdb"), // 数据库会自动创建这个目录文件
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = db.Close() }()

	// 把这段注释后, 就获取不到值了
	//if err := db.Update(
	//	func(tx *nutsdb.Tx) error {
	//		key := []byte("name1")
	//		val := []byte("val12")
	//		bucket := "bucket1"
	//		if err := tx.Put(bucket, key, val, 0); err != nil {
	//			return err
	//		}
	//		return nil
	//	}); err != nil {
	//	log.Fatal(err)
	//}

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			key := []byte("name1")
			bucket := "bucket1"
			if e, err := tx.Get(bucket, key); err != nil {
				return err
			} else {
				fmt.Println(string(e.Value)) // "val1-modify"
			}
			return nil
		}); err != nil {
		log.Println(err)
	}
}
