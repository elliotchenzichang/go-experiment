package _go

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
}

func TestMapDelete(t *testing.T) {
	var data = map[string]string{
		"name": "Elliot",
	}
	fmt.Println(data["name"])

	delete(data, "name")

	fmt.Println(data["name"])

}