package collections

import (
	"errors"
	"fmt"
	"strings"
)

type Node[T any] struct {
	Next  *Node[T]
	Prev  *Node[T]
	Value T
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.Value)
}

/*
Implemented as a Doubly Linked List

Fast operations on front and back, slow iteration through.
*/
type Deque[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{}
}

func (l *Deque[T]) Length() int {
	return l.Size
}

// Visualizes the deque as a simple string,
// for debugging.
func (l *Deque[T]) Display() string {

	res := make([]string, l.Size)

	curr := l.Head

	for i := 0; curr != nil; i++ {
		res[i] = curr.String()
		curr = curr.Next
	}

	return "[" + strings.Join(res, " ") + "]"

}

func (l *Deque[T]) Append(value T) {

	if l.Size == 0 {

		node := &Node[T]{
			Next:  nil,
			Prev:  nil,
			Value: value,
		}

		l.Head = node
		l.Tail = node

	} else {

		node := &Node[T]{
			Next:  nil,
			Prev:  l.Tail,
			Value: value,
		}

		l.Tail.Next = node
		l.Tail = node

	}

	l.Size++

}

func (l *Deque[T]) AppendLeft(value T) {

	if l.Size == 0 {

		node := &Node[T]{
			Next:  nil,
			Prev:  nil,
			Value: value,
		}

		l.Head = node
		l.Tail = node

	} else {

		node := &Node[T]{
			Next:  l.Head,
			Prev:  nil,
			Value: value,
		}

		l.Head.Prev = node
		l.Head = node

	}

	l.Size++

}

func (l *Deque[T]) Extend(values []T) {

	for _, v := range values {
		l.Append(v)
	}

}

func (l *Deque[T]) ExtendLeft(values []T) {

	for _, v := range values {
		l.AppendLeft(v)
	}

}

func (l *Deque[T]) Pop() (T, error) {

	var result T

	if l.Size == 0 {
		return result, errors.New("cannot pop an empty doubly linked list")
	}

	result = l.Tail.Value
	l.Tail = l.Tail.Prev
	l.Tail.Next = nil
	l.Size--

	if l.Size == 0 {
		l.Head = nil
	}

	return result, nil

}

func (l *Deque[T]) PopLeft() (T, error) {

	var result T

	if l.Size == 0 {
		return result, errors.New("cannot pop an empty doubly linked list")
	}

	result = l.Head.Value
	l.Head = l.Head.Next
	l.Head.Prev = nil
	l.Size--

	if l.Size == 0 {
		l.Tail = nil
	}

	return result, nil

}

/*
1 2 3 (Size = 3)

Insert @ 2
*/
func (l *Deque[T]) Insert(value T, index int) {

	if index > l.Size-1 {
		l.Append(value)
		return
	}

	if index == 0 {
		l.AppendLeft(value)
		return
	}

	curr := l.Head

	for i := 1; curr != nil; i++ {

		if i == index {
			break
		}
		curr = curr.Next

	}

	node := &Node[T]{
		Next:  curr.Next,
		Prev:  curr,
		Value: value,
	}

	node.Prev.Next = node
	node.Next.Prev = node

	l.Size++

}

func (l *Deque[T]) Reverse() {

	if l.Size <= 1 {
		return
	}

	var temp *Node[T]
	curr := l.Head

	l.Tail = l.Head

	for curr != nil {
		temp = curr.Prev
		curr.Prev = curr.Next
		curr.Next = temp
		curr = curr.Prev
	}

	l.Head = temp.Prev

}

func (l *Deque[T]) Values() []T {

	if l.Size == 0 {
		var res []T
		return res
	}

	res := make([]T, l.Size)
	curr := l.Head

	for i := 0; curr != nil; i++ {
		res[i] = curr.Value
		curr = curr.Next
	}

	return res
}

func (l *Deque[T]) Copy() *Deque[T] {

	_l := NewDeque[T]()
	curr := l.Head

	for curr != nil {
		_l.Append(curr.Value)
		curr = curr.Next
	}

	return _l

}

func DequeIndex[T comparable](d *Deque[T], target T) (int, error) {

	var counter int

	curr := d.Head
	for curr != nil {
		if curr.Value == target {
			return counter, nil
		}
	}

	return counter, fmt.Errorf("value %v not found in deque", target)

}

func DequeCount[T comparable](d *Deque[T], target T) int {

	counter := 0
	curr := d.Head
	for curr != nil {
		if curr.Value == target {
			counter++
		}
	}

	return counter
}
