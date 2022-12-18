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
