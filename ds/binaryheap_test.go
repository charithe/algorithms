package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapInsertion(t *testing.T) {
	t.Parallel()
	cmp := func(a interface{}, b interface{}) int {
		return a.(int) - b.(int)
	}

	bh := NewBinaryHeap(cmp, 4)
	err := bh.Insert(10)
	assert.NoError(t, err)
	assert.Equal(t, 10, bh.Peek())
	assert.Equal(t, 1, bh.Size())

	err = bh.Insert(20)
	assert.NoError(t, err)
	assert.Equal(t, 20, bh.Peek())
	assert.Equal(t, 2, bh.Size())

	err = bh.Insert(15)
	assert.NoError(t, err)
	assert.Equal(t, 20, bh.Peek())
	assert.Equal(t, 3, bh.Size())

	err = bh.Insert(40)
	assert.NoError(t, err)
	assert.Equal(t, 40, bh.Peek())
	assert.Equal(t, 4, bh.Size())

	err = bh.Insert(50)
	assert.Error(t, err)
}

func TestHeapRemoval(t *testing.T) {
	t.Parallel()
	cmp := func(a interface{}, b interface{}) int {
		return a.(int) - b.(int)
	}

	bh := NewBinaryHeap(cmp, 4)
	err := bh.Insert(10, 20, 15, 40)
	assert.NoError(t, err)
	assert.Equal(t, 4, bh.Size())

	item, err := bh.Remove()
	assert.NoError(t, err)
	assert.Equal(t, 40, item)

	item, err = bh.Remove()
	assert.NoError(t, err)
	assert.Equal(t, 20, item)

	item, err = bh.Remove()
	assert.NoError(t, err)
	assert.Equal(t, 15, item)

	item, err = bh.Remove()
	assert.NoError(t, err)
	assert.Equal(t, 10, item)

	_, err = bh.Remove()
	assert.Error(t, err)
}
