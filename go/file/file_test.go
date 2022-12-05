package file

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"testing"
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
