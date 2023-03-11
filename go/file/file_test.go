package file

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

var path = "test_file.txt"

func TestPrepareData(t *testing.T) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	for i := 1; i <= 20000000; i++ {
		write.WriteString(fmt.Sprint(i))
	}
	write.Flush()
}

func BenchmarkFile_ReadAt(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	skipLength := 512
	var index int64 = 0
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	for {
		buf := make([]byte, skipLength)
		_, err := file.ReadAt(buf, index)
		if err != nil {
			return
		}
		index += int64(skipLength)
	}
}

func BenchmarkFile_Read1(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	skipLength := 512
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(file)
	for {
		buf := make([]byte, skipLength)
		n, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		if n == 0 {
			break
		}
	}
}

func TestRead(t *testing.T) {
	fd, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	data := []byte("aaaaaaaa")
	write, err := fd.Write(data)
	if err != nil {
		return
	}
	if write != len(data) {
		t.Fatal("write data length unexpected")
	}

	fd2, err := os.OpenFile("test.txt", os.O_RDWR, os.ModePerm)
	readData := make([]byte, 4096)
	read, err := fd2.Read(readData)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(read)
}

func TestBufioReader(t *testing.T) {
	fd, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	dataSize := 5000
	data := make([]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = byte(rand.Intn(100))
	}
	_, err = fd.Write(data)
	if err != nil {
		return
	}

	fd2, err := os.OpenFile("test.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	reader := bufio.NewReader(fd2)
	block1 := make([]byte, 3000)
	block2 := make([]byte, 2000)

	read, err := reader.Read(block1)
	if err != nil {
		return
	}
	fmt.Println(read)
	full, err := io.ReadFull(reader, block2)
	if err != nil {
		return
	}
	fmt.Println(full)
}

func TestBufio_ReadWriter(t *testing.T) {
	readFd, _ := os.OpenFile("read.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	writerFd, _ := os.OpenFile("writer.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	reader := bufio.NewReader(readFd)
	writer := bufio.NewWriter(writerFd)
	rw := bufio.NewReadWriter(reader, writer)
	r := make([]byte, 1024)
	w := make([]byte, 1024)
	rw.Read(r)
	rw.Write(w)
}

var task = func(numberOfGoroutines int, t *testing.T) {
	f, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	bytes := []byte(strings.Repeat("a", 4*1024))
	rw := &sync.RWMutex{}
	wg := &sync.WaitGroup{}
	start := time.Now().UnixNano()
	for i := 0; i < numberOfGoroutines; i++ {
		writeTimes := 1024 / numberOfGoroutines
		go func() {
			wg.Add(1)
			defer wg.Done()
			for i := 0; i < writeTimes; i++ {
				rw.Lock()
				f.Write(bytes)
				rw.Unlock()
			}
		}()
	}
	wg.Wait()
	final := time.Now().UnixNano()
	t.Logf("try number of goroutine: %d, and use time is %d", numberOfGoroutines, final-start)
	os.Remove("test.txt")
}

func TestConcurrentReadByMultiGoroutine8(t *testing.T) {
	task(8, t)
}
func TestConcurrentReadByMultiGoroutine16(t *testing.T) {
	task(16, t)
}
func TestConcurrentReadByMultiGoroutine32(t *testing.T) {
	task(32, t)
}

var taskOfChannelTest = func(numberOfGoroutines int, t *testing.T) {
	wg := &sync.WaitGroup{}

	bytes := []byte(strings.Repeat("a", 4*1024))
	type param struct {
		r     chan struct{}
		bytes []byte
	}
	c := make(chan *param, 1)
	f, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	go func() {
		for {
			select {
			case param := <-c:
				f.Write(param.bytes)
				param.r <- struct{}{}
			}
		}

	}()
	start := time.Now().UnixNano()
	for i := 0; i < numberOfGoroutines; i++ {
		writeTimes := 1024 / numberOfGoroutines
		go func() {
			wg.Add(1)
			defer wg.Done()
			r := make(chan struct{})
			for i := 0; i < writeTimes; i++ {
				p := &param{
					r:     r,
					bytes: bytes,
				}
				c <- p
				select {
				case <-p.r:

				}
			}
		}()
	}
	wg.Wait()
	final := time.Now().UnixNano()
	t.Logf("try number of goroutine write via channel: %d, and use time is %d", numberOfGoroutines, final-start)
	os.Remove("test.txt")
}

func TestWriteDataBySingleGoroutineViaChannel8(t *testing.T) {
	taskOfChannelTest(8, t)
}

func TestWriteDataBySingleGoroutineViaChannel16(t *testing.T) {
	taskOfChannelTest(16, t)
}

func TestWriteDataBySingleGoroutineViaChannel32(t *testing.T) {
	taskOfChannelTest(32, t)
}

var taskOfChannelTestForBM = func(numberOfGoroutines int, t *testing.B) {
	wg := &sync.WaitGroup{}

	bytes := []byte(strings.Repeat("a", 4*1024))
	type param struct {
		r     chan struct{}
		bytes []byte
	}
	c := make(chan *param, 1)
	f, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	go func() {
		for {
			select {
			case param := <-c:
				f.Write(param.bytes)
				f.Sync()
				param.r <- struct{}{}
			}
		}

	}()
	for i := 0; i < numberOfGoroutines; i++ {
		writeTimes := 3
		go func() {
			wg.Add(1)
			defer wg.Done()
			r := make(chan struct{})
			for i := 0; i < writeTimes; i++ {
				p := &param{
					r:     r,
					bytes: bytes,
				}
				c <- p
				select {
				case <-p.r:

				}
			}
		}()
	}
	wg.Wait()
}

func BenchmarkWriteDataBySingleGoroutineViaChannel(b *testing.B) {
	for _, number := range []int{8, 16, 32} {
		b.Run(fmt.Sprintf("test for %d goroutines", number), func(b *testing.B) {
			b.StartTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				taskOfChannelTestForBM(number, b)
			}
		})
	}
}

var taskForBM = func(numberOfGoroutines int, b *testing.B) {
	f, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	bytes := []byte(strings.Repeat("a", 4*1024))
	rw := &sync.RWMutex{}
	wg := &sync.WaitGroup{}
	for i := 0; i < numberOfGoroutines; i++ {
		writeTimes := 3
		go func() {
			wg.Add(1)
			defer wg.Done()
			for i := 0; i < writeTimes; i++ {
				rw.Lock()
				f.Write(bytes)
				f.Sync()
				rw.Unlock()
			}
		}()
	}
	wg.Wait()
	os.Remove("test.txt")
}

func BenchmarkWriteDataByMultiGoroutines(b *testing.B) {
	for _, number := range []int{8, 16, 32} {
		b.Run(fmt.Sprintf("test for %d goroutines", number), func(b *testing.B) {
			b.StartTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				taskForBM(number, b)
			}
		})
	}
}
