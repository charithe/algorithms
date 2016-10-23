package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueueAndDequeue(t *testing.T) {
	var err error
	q := NewCircularQueue(2)
	err = q.Enqueue("a")
	assert.NoError(t, err)
	err = q.Enqueue("b")
	assert.NoError(t, err)
	err = q.Enqueue("c")
	assert.Error(t, err)

	var item interface{}
	item, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, item, "a")

	item, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, item, "b")

	item, err = q.Dequeue()
	assert.Error(t, err)
}
