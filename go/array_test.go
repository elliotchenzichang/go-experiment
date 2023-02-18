package _go

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arrays := []int{1, 2, 3, 4}
	fmt.Println(arrays[0:3])
}

type LinkedList struct {
	head *Node
	tail *Node
	node *Node
}

type Node struct {
	val  *Item
	next *Node
}

type Item struct {
	val int
}

func NewLinkedList() *LinkedList {
	l := &LinkedList{}
	l.head = &Node{}
	l.tail = l.head
	return l
}

func (l *LinkedList) Add(node *Node) {
	l.tail.next = node
	l.tail = l.tail.next
}

func (l *LinkedList) AddNodeInside() {
	node := &Node{
		val:  &Item{},
		next: nil,
	}
	l.tail.next = node
	l.tail = l.tail.next
}

func BenchmarkLinkList(b *testing.B) {
	l := NewLinkedList()
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		node := &Node{val: &Item{val: 1}}
		l.Add(node)
	}
}

func BenchmarkLinkListAddInside(b *testing.B) {
	l := NewLinkedList()
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.AddNodeInside()
	}
}

func BenchmarkAppend(b *testing.B) {
	b.ReportAllocs()
	b.StartTimer()
	var slice []Item
	for i := 0; i < b.N; i++ {
		item := Item{val: 1}
		slice = append(slice, item)
	}
}

func BenchmarkAppendPointer(b *testing.B) {
	b.ReportAllocs()
	b.StartTimer()
	var slice []*Item
	for i := 0; i < b.N; i++ {
		item := &Item{val: 1}
		slice = append(slice, item)
	}
}

func BenchmarkSlickGrow(b *testing.B) {
	// 要测试的切片长度
	var lengths = []int{1000, 10 * 1000, 100 * 1000, 1000 * 1000}
	for _, length := range lengths {
		// 直接申请空间的切片 性能测试
		nameOfNotGrowBM := fmt.Sprintf("test_slice_not_grow_%d", length)
		b.Run(nameOfNotGrowBM, func(b *testing.B) {
			b.ReportAllocs()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				value := 1
				slice := make([]int, length)
				for i := 0; i < length; i++ {
					slice = append(slice, value)
				}
			}
		})
		// 从一开始就不申请空间，一路append的切片 性能测试
		nameOfGrowBM := fmt.Sprintf("test_slice_grow_%d", length)
		b.Run(nameOfGrowBM, func(b *testing.B) {
			b.ReportAllocs()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				value := 1
				var slice []int
				for i := 0; i < length; i++ {
					slice = append(slice, value)
				}
			}
		})
	}
}

func TestSliceBaseUsage(t *testing.T) {
	var slice []int
	slice = append(slice, 1, 2, 3)
}
