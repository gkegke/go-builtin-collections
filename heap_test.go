package collections

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {

	H := NewHeap[int](func(a, b int) bool { return a < b })

	H.Push(100)
	H.Push(1)
	H.Push(18)
	H.Push(1000)

	fmt.Println(H.Heap)
	a, _ := H.Pop()
	fmt.Println(H.Heap)
	b, _ := H.Pop()
	fmt.Println(H.Heap)
	c, _ := H.Pop()
	fmt.Println(H.Heap)

	tests := []struct {
		Expected int
		Got      int
	}{
		{1, a},
		{18, b},
		{100, c},
	}

	for _, test := range tests {

		assert.Equal(
			t,
			test.Expected,
			test.Got,
			"",
		)

	}

}

func TestMaxHeap(t *testing.T) {

	H := NewHeap[int](func(a, b int) bool { return a > b })

	H.Push(100)
	H.Push(1)
	H.Push(18)
	H.Push(1000)

	fmt.Println(H.Heap)
	a, _ := H.Pop()
	fmt.Println(H.Heap)
	b, _ := H.Pop()
	fmt.Println(H.Heap)
	c, _ := H.Pop()
	fmt.Println(H.Heap)

	tests := []struct {
		Expected int
		Got      int
	}{
		{1000, a},
		{100, b},
		{18, c},
	}

	for _, test := range tests {

		assert.Equal(
			t,
			test.Expected,
			test.Got,
			"",
		)

	}

}

func TestMinHeap2(t *testing.T) {

	H := NewHeap[int](func(a, b int) bool { return a < b })

	for _, d := range []int{2, 33, 4, 15, 7, 9, 10, 1} {
		H.Push(d)
	}

	fmt.Println(H.Heap)
	a, _ := H.Pop()
	fmt.Println(H.Heap)
	b, _ := H.Pop()
	fmt.Println(H.Heap)
	c, _ := H.Pop()
	fmt.Println(H.Heap)

	tests := []struct {
		Expected int
		Got      int
	}{
		{1, a},
		{2, b},
		{4, c},
	}

	for _, test := range tests {

		assert.Equal(
			t,
			test.Expected,
			test.Got,
			"",
		)

	}

}
