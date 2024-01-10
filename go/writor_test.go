package _go

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"testing"
)

var writeStr = "Hello World"
var readBytes = make([]byte, 11)

type ReadWrite interface {
	Read()
	Write()
}

type TestConfig struct {
	flowControlThreshold int
}

func initTestResources() (rw *RWLockWriter, vw *VersionControlWriter, err error) {
	fd, errForRw := os.OpenFile("test-1.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if errForRw != nil {
		return nil, nil, errForRw
	}
	fd2, errForVersion := os.OpenFile("test-2.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if errForVersion != nil {
		return nil, nil, errForVersion
	}
	rw = &RWLockWriter{
		lock: sync.RWMutex{},
		fd:   fd,
	}
	vw = &VersionControlWriter{
		version: atomic.Value{},
		fd:      fd2,
	}
	vw.version.Store(0)
	rw.Write()
	vw.Write()
	return rw, vw, nil
}

func closeTestResources(fw *RWLockWriter, vw *VersionControlWriter) error {
	err := fw.fd.Close()
	if err != nil {
		return err
	}
	err = vw.fd.Close()
	if err != nil {
		return err
	}
	err = os.Remove("test-1.txt")
	if err != nil {
		return err
	}
	err = os.Remove("test-2.txt")
	if err != nil {
		return err
	}
	return nil
}

func BenchmarkWriterPerformance(b *testing.B) {
	for _, config := range []TestConfig{
		{
			flowControlThreshold: 20,
		},
		{
			flowControlThreshold: 50,
		},
		{
			flowControlThreshold: 80,
		},
	} {
		rw, vw, err := initTestResources()
		if err != nil {
			b.Error(err)
		}
		b.Run(fmt.Sprintf("test the performance for RWLock for flowcontorl %d", config.flowControlThreshold), func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					flowControl(config.flowControlThreshold, rw)
				}
			})

		})
		b.Run(fmt.Sprintf("test the performance for version control flowcontorl %d", config.flowControlThreshold), func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					flowControl(config.flowControlThreshold, vw)
				}
			})
		})
		err = closeTestResources(rw, vw)
		if err != nil {
			b.Error(err)
		}
	}

}

type VersionControlWriter struct {
	version atomic.Value
	fd      *os.File
}

func flowControl(threshold int, rw ReadWrite) {
	seed := rand.Intn(100)
	if seed < threshold {
		rw.Write()
	} else {
		rw.Read()
	}
}

func (vw *VersionControlWriter) Write() {
	vw.increase()
	vw.fd.Write([]byte(writeStr))
}

func (vw *VersionControlWriter) Read() {
	vw.increase()
	vw.fd.ReadAt(readBytes, 0)
}

func (vw *VersionControlWriter) increase() {
	v := vw.version.Load().(int)
	v++
	vw.version.Store(v)
}

type RWLockWriter struct {
	lock sync.RWMutex
	fd   *os.File
}

func (rw *RWLockWriter) Write() {
	rw.lock.Lock()
	defer rw.lock.Unlock()
	rw.fd.Write([]byte(writeStr))
}

func (rw *RWLockWriter) Read() {
	rw.lock.RLock()
	defer rw.lock.RUnlock()
	rw.fd.ReadAt(readBytes, 0)
}
