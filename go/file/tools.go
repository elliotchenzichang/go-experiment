package file

import (
	"os"
	"testing"
)

func RunFileTest(t *testing.T, resourceCreator func() (*os.File, error), resourceDestroy func(f *os.File), callback func(f *os.File) error) {
	rw, err := resourceCreator()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		resourceDestroy(rw)
	}()
	err = callback(rw)
	if err != nil {
		t.Fatal(err)
	}
}
