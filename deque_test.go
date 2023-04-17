package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDequeAppend(t *testing.T) {

	d := NewDeque[int]()

	assert.Equal(
		t,
		0,
		d.Length(),
		"",
	)

	d.Append(1)

	tests := []struct {
		Expected int
		Got      int
	}{
		{1, d.Length()},
		{1, d.Head.Value},
		{1, d.Tail.Value},
	}

	for _, test := range tests {

		assert.Equal(
			t,
			test.Expected,
			test.Got,
			"",
		)
	}

	d.Append(3)
	d.Append(5)

	tests = []struct {
		Expected int
		Got      int
	}{
		{3, d.Length()},
		{1, d.Head.Value},
		{5, d.Tail.Value},
	}

	for _, test := range tests {

		assert.Equal(
			t,
			test.Expected,
			test.Got,
			"",
		)
	}

	assert.Equal(
		t,
		[]int{1, 3, 5},
		d.Values(),
		"",
	)

}

func TestDequeAppendLeft(t *testing.T) {

	d := NewDeque[int]()

	d.AppendLeft(1)

	tests := []struct {
		Expected int
		Got      int
	}{
		{1, d.Length()},
		{1, d.Head.Value},
		{1, d.Tail.Value},
	}

	for _, test := range tests {

		assert.Equal(
			t,
			test.Expected,
			test.Got,
			"",
		)
	}

	d.AppendLeft(3)
	d.AppendLeft(5)

	tests = []struct {
		Expected int
		Got      int
	}{
		{3, d.Length()},
		{5, d.Head.Value},
		{1, d.Tail.Value},
	}

	for _, test := range tests {

		assert.Equal(
			t,
			test.Expected,
			test.Got,
			"",
		)
	}

	assert.Equal(
		t,
		[]int{5, 3, 1},
		d.Values(),
		"",
	)

}

func TestDequeExtend(t *testing.T) {

	d := NewDeque[int]()

	d.Extend([]int{1, 2, 3, 4})

	tests := []struct {
		Expected int
		Got      int
	}{
		{4, d.Length()},
		{1, d.Head.Value},
		{4, d.Tail.Value},
	}

	for _, test := range tests {

		assert.Equal(
			t,
			test.Expected,
			test.Got,
			"",
		)
	}

	assert.Equal(
		t,
		[]int{1, 2, 3, 4},
		d.Values(),
		"",
	)

}

func TestDequeExtendLeft(t *testing.T) {

	d := NewDeque[int]()

	d.ExtendLeft([]int{1, 2, 3, 4})

	tests := []struct {
		Expected int
		Got      int
	}{
		{4, d.Length()},
		{4, d.Head.Value},
		{1, d.Tail.Value},
	}

	for _, test := range tests {

		assert.Equal(
			t,
			test.Expected,
			test.Got,
			"",
		)
	}

	assert.Equal(
		t,
		[]int{4, 3, 2, 1},
		d.Values(),
		"",
	)

}

func TestDequePop(t *testing.T) {

	d := NewDeque[int]()

	d.Extend([]int{1, 2, 3, 4})

	x, _ := d.Pop()

	assert.Equal(t, 4, x, "")

	d.Extend([]int{5, 6, 7, 8})

	_, _ = d.Pop()
	_, _ = d.Pop()
	_, _ = d.Pop()

	tests := []struct {
		Expected int
		Got      int
	}{
		{4, d.Length()},
		{1, d.Head.Value},
		{5, d.Tail.Value},
	}

	for _, test := range tests {

		assert.Equal(
			t,
			test.Expected,
			test.Got,
			"",
		)
	}

	assert.Equal(
		t,
		[]int{1, 2, 3, 5},
		d.Values(),
		"",
	)

}

func TestDequePopLeft(t *testing.T) {

	d := NewDeque[int]()

	d.Extend([]int{1, 2, 3, 4})

	x, _ := d.PopLeft()

	assert.Equal(t, 1, x, "")

	d.Extend([]int{5, 6, 7, 8})

	_, _ = d.PopLeft()
	_, _ = d.PopLeft()
	_, _ = d.PopLeft()

	tests := []struct {
		Expected int
		Got      int
	}{
		{4, d.Length()},
		{5, d.Head.Value},
		{8, d.Tail.Value},
	}

	for _, test := range tests {

		assert.Equal(
			t,
			test.Expected,
			test.Got,
			"",
		)
	}

	assert.Equal(
		t,
		[]int{5, 6, 7, 8},
		d.Values(),
		"",
	)

}

func TestDequeInsert(t *testing.T) {

	d := NewDeque[int]()

	d.Insert(1, 5)

	assert.Equal(
		t,
		[]int{1},
		d.Values(),
		"",
	)

	d.Extend([]int{1, 2, 3, 4})

	d.Insert(100, 2)

	tests := []struct {
		Expected int
		Got      int
	}{
		{6, d.Length()},
		{1, d.Head.Value},
		{4, d.Tail.Value},
	}

	for _, test := range tests {

		assert.Equal(
			t,
			test.Expected,
			test.Got,
			"",
		)

	}

	assert.Equal(
		t,
		[]int{1, 1, 100, 2, 3, 4},
		d.Values(),
		"",
	)

}

func TestDequeReverse(t *testing.T) {

	d := NewDeque[int]()

	d.Extend([]int{1, 2, 3, 4})

	d.Reverse()

	tests := []struct {
		Expected int
		Got      int
	}{
		{4, d.Length()},
		{4, d.Head.Value},
		{1, d.Tail.Value},
	}

	for _, test := range tests {

		assert.Equal(
			t,
			test.Expected,
			test.Got,
			"",
		)

	}

	assert.Equal(
		t,
		[]int{4, 3, 2, 1},
		d.Values(),
		"",
	)

}

func TestValues(t *testing.T) {

	d := NewDeque[int]()

	d.Extend([]int{1, 2, 3, 4})

	tests := []struct {
		Expected []int
		Got      []int
	}{
		{[]int{1, 2, 3, 4}, d.Values()},
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
