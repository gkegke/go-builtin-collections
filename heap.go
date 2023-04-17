package collections

import (
	"errors"
)

/*

Heap Order Property:
 - In a minHeap, for every node x with parent p,
   the key in p is smaller than or equal to the key in x.

*/

type Heap[T any] struct {
	Heap []T
	comp func(a, b T) bool
}

func NewHeap[T any](comp func(a, b T) bool) *Heap[T] {
	return &Heap[T]{comp: comp}
}

func (h *Heap[T]) Swap(i, j int) {
	h.Heap[i], h.Heap[j] = h.Heap[j], h.Heap[i]
}

// Move value up in the tree to maintain heap order property.
func (h *Heap[T]) percUp(i int) {

	p := parent(i)

	for {

		if !(i > 0 && h.comp(h.Heap[i], h.Heap[p])) {
			break
		}

		h.Swap(i, p)
		i = p
		p = int((p - 1) / 2)

	}

}

// Move value down in the tree to maintain heap order property.
func (h *Heap[T]) percDown(i int) {

	hlen := len(h.Heap)

	mc := 0

	for {

		// get min child index
		l := left(i)

		if l >= hlen || l < 0 {
			break
		}

		r := right(i)

		if r < hlen && h.comp(h.Heap[r], h.Heap[l]) {
			mc = r
		} else {
			mc = l
		}

		// if child is smaller, swap
		if !(h.comp(h.Heap[mc], h.Heap[i])) {
			break
		}

		h.Swap(i, mc)
		i = mc

	}

}

// Returns the prioritised value in the heap.
// e.g. in a min-heap, it'll return the minimum value.
func (h *Heap[T]) Pop() (T, error) {

	/*

	  1. swap the initial value for the last value
	  2. remove the last value in the heap
	  3. reorder the heap to satisfy ordering conditions

	*/

	var result T

	if len(h.Heap) == 0 {
		return result, errors.New("cannot pop an empty heap")
	}

	result = h.Heap[0]

	hlen := len(h.Heap) - 1

	h.Swap(0, hlen)

	h.Heap = h.Heap[0:hlen]

	h.percDown(0)

	return result, nil

}

// Add new values to the heap
func (h *Heap[T]) Push(v T) {

	h.Heap = append(h.Heap, v)
	h.percUp(len(h.Heap) - 1)

}

// helper functions
func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return (i * 2) + 1 }
func right(i int) int  { return (i * 2) + 2 }
