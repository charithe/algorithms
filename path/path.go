package path

import (
	"bytes"
	"fmt"

	"github.com/fatih/color"
)

type PathFinder func(*Map, *Coordinate, *Coordinate) Path

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

func NewCoord(x int, y int) *Coordinate {
	return &Coordinate{
		X: x,
		Y: y,
	}
}

type weightedCoord struct {
	coord  *Coordinate
	weight int
}

func weightedCoordComparer(a interface{}, b interface{}) int {
	ac := a.(*weightedCoord)
	bc := b.(*weightedCoord)
	return bc.weight - ac.weight
}

type Path []*Coordinate

func (p Path) String() string {
	if len(p) == 0 {
		return "No Path"
	}

	var buffer bytes.Buffer
	for i := len(p) - 1; i > 0; i-- {
		buffer.WriteString(p[i].String())
		buffer.WriteString(" -> ")
	}
	buffer.WriteString(p[0].String())
	return buffer.String()
}

type Obstacle struct {
	From Coordinate
	To   Coordinate
}

func NewObstacle(x1 int, y1 int, x2 int, y2 int) *Obstacle {
	return &Obstacle{
		From: Coordinate{X: x1, Y: y1},
		To:   Coordinate{X: x2, Y: y2},
	}
}

type Cell int

const (
	Free Cell = iota
	Blocked
	PathPoint
)

type Map struct {
	cells [][]Cell
	size  int
}

func NewMap(size int, obstacles ...*Obstacle) *Map {
	cells := make([][]Cell, size)
	for i := 0; i < size; i++ {
		cells[i] = make([]Cell, size)
	}

	for _, obstacle := range obstacles {
		xDist := obstacle.To.X - obstacle.From.X
		yDist := obstacle.To.Y - obstacle.From.Y
		for x := 0; x <= xDist; x++ {
			for y := 0; y <= yDist; y++ {
				cells[obstacle.From.Y+y][obstacle.From.X+x] = Blocked
			}
		}
	}

	return &Map{
		cells: cells,
		size:  size,
	}
}

func (m *Map) Draw() {
	blocked := color.New(color.FgBlue)
	pathPoint := color.New(color.FgRed)

	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			switch m.cells[i][j] {
			case Free:
				fmt.Printf(". ")
			case Blocked:
				blocked.Printf("# ")
			case PathPoint:
				pathPoint.Printf("* ")
			}
		}
		fmt.Print("\n")
	}
}

func (m *Map) DrawPath(path Path) {
	newCells := make([][]Cell, m.size)
	for i := 0; i < m.size; i++ {
		newCells[i] = make([]Cell, m.size)
		for j := 0; j < m.size; j++ {
			newCells[i][j] = m.cells[i][j]
		}
	}

	for _, pathPoint := range path {
		newCells[pathPoint.Y][pathPoint.X] = PathPoint
	}

	tmp := &Map{cells: newCells, size: m.size}
	tmp.Draw()
}

func (m *Map) IsFree(coord *Coordinate) bool {
	return m.cells[coord.Y][coord.X] == Free
}

func (m *Map) GetNeighbours(coord *Coordinate) []*Coordinate {
	var neighbours []*Coordinate

	for xd := -1; xd <= 1; xd++ {
		for yd := -1; yd <= 1; yd++ {
			if newX := coord.X + xd; newX >= 0 && newX < m.size {
				if newY := coord.Y + yd; newY >= 0 && newY < m.size {
					newCoord := &Coordinate{X: newX, Y: newY}
					if (*newCoord) != (*coord) {
						neighbours = append(neighbours, newCoord)
					}
				}
			}
		}
	}

	return neighbours
}
