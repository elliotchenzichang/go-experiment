package _go

import (
	"fmt"
	"testing"
)

func TestMapDelete(t *testing.T) {
	var data = map[string]string{
		"name": "Elliot",
	}
	fmt.Println(data["name"])

	delete(data, "name")

	fmt.Println(data["name"])

}

func BenchmarkPerformanceOfGrowthMapOrNot(b *testing.B) {
	lengths := []int{1000, 10000, 100000}
	for _, length := range lengths {
		b.Run(fmt.Sprintf("test the performance of the ungrow map with the length of %d", length), func(b *testing.B) {
			m := make(map[string]string, length)
			for i := 0; i < length; i++ {
				key := fmt.Sprintf("test_key_%d", i)
				value := fmt.Sprintf("value_%d", i)
				m[key] = value
			}
		})
		b.Run(fmt.Sprintf("test the performance of the grow map with the length of %d", length), func(b *testing.B) {
			m := make(map[string]string)
			for i := 0; i < length; i++ {
				key := fmt.Sprintf("test_key_%d", i)
				value := fmt.Sprintf("value_%d", i)
				m[key] = value
			}
		})
	}
}
