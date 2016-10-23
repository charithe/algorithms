package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNeighboursInMiddle(t *testing.T) {
	m := NewMap(10)
	neighbours := m.GetNeighbours(&Coordinate{X: 2, Y: 2})

	assert.Len(t, neighbours, 8)
	assert.Contains(t, neighbours, &Coordinate{X: 1, Y: 1})
	assert.Contains(t, neighbours, &Coordinate{X: 1, Y: 2})
	assert.Contains(t, neighbours, &Coordinate{X: 1, Y: 3})
	assert.Contains(t, neighbours, &Coordinate{X: 2, Y: 1})
	assert.Contains(t, neighbours, &Coordinate{X: 2, Y: 3})
	assert.Contains(t, neighbours, &Coordinate{X: 3, Y: 1})
	assert.Contains(t, neighbours, &Coordinate{X: 3, Y: 2})
	assert.Contains(t, neighbours, &Coordinate{X: 3, Y: 3})
}

func TestGetNeighboursInTopEdge(t *testing.T) {
	m := NewMap(10)
	neighbours := m.GetNeighbours(&Coordinate{X: 0, Y: 2})

	assert.Len(t, neighbours, 5)
	assert.Contains(t, neighbours, &Coordinate{X: 0, Y: 1})
	assert.Contains(t, neighbours, &Coordinate{X: 0, Y: 3})
	assert.Contains(t, neighbours, &Coordinate{X: 1, Y: 1})
	assert.Contains(t, neighbours, &Coordinate{X: 1, Y: 2})
	assert.Contains(t, neighbours, &Coordinate{X: 1, Y: 3})
}

func TestGetNeighboursInBottomEdge(t *testing.T) {
	m := NewMap(10)
	neighbours := m.GetNeighbours(&Coordinate{X: 9, Y: 3})

	assert.Len(t, neighbours, 5)
	assert.Contains(t, neighbours, &Coordinate{X: 9, Y: 2})
	assert.Contains(t, neighbours, &Coordinate{X: 9, Y: 4})
	assert.Contains(t, neighbours, &Coordinate{X: 8, Y: 2})
	assert.Contains(t, neighbours, &Coordinate{X: 8, Y: 3})
	assert.Contains(t, neighbours, &Coordinate{X: 8, Y: 4})
}

func TestGetNeighboursInTopCorner(t *testing.T) {
	m := NewMap(10)
	neighbours := m.GetNeighbours(&Coordinate{X: 0, Y: 0})

	assert.Len(t, neighbours, 3)
	assert.Contains(t, neighbours, &Coordinate{X: 0, Y: 1})
	assert.Contains(t, neighbours, &Coordinate{X: 1, Y: 0})
	assert.Contains(t, neighbours, &Coordinate{X: 1, Y: 1})
}

func TestGetNeighboursInBottomCorner(t *testing.T) {
	m := NewMap(10)
	neighbours := m.GetNeighbours(&Coordinate{X: 9, Y: 9})

	assert.Len(t, neighbours, 3)
	assert.Contains(t, neighbours, &Coordinate{X: 9, Y: 8})
	assert.Contains(t, neighbours, &Coordinate{X: 8, Y: 9})
	assert.Contains(t, neighbours, &Coordinate{X: 8, Y: 8})
}
