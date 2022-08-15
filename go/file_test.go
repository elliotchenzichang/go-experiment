package _go

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

var path = "data/test_file.txt"

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
	skipLength := 50
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
