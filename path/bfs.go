package path

import (
	"math"

	"github.com/charithe/algorithms/ds"
)

func BreadthFirst(m *Map, from *Coordinate, to *Coordinate) Path {
	var path []*Coordinate
	cameFrom := make(map[Coordinate]*Coordinate)
	frontier := ds.NewCircularQueue(m.size * m.size)

	frontier.Enqueue(from)
	cameFrom[*from] = from

	for {
		temp, err := frontier.Dequeue()
		if err != nil {
			break
		}

		coord := temp.(*Coordinate)

		if (*coord) == (*to) {
			pathPoint := coord
			for {
				path = append(path, pathPoint)
				if (*pathPoint) == (*from) {
					break
				}
				pathPoint = cameFrom[*pathPoint]
			}
			break
		}

		neighbours := m.GetNeighbours(coord)
		for _, neighbour := range neighbours {
			if _, ok := cameFrom[*neighbour]; !ok && m.IsFree(neighbour) {
				frontier.Enqueue(neighbour)
				cameFrom[*neighbour] = coord
			}
		}
	}

	return path
}

func GreedyBreadthFirst(m *Map, from *Coordinate, to *Coordinate) Path {
	var path []*Coordinate
	cameFrom := make(map[Coordinate]*Coordinate)
	frontier := ds.NewBinaryHeap(weightedCoordComparer, m.size*m.size)

	frontier.Insert(&weightedCoord{coord: from, weight: 0})
	cameFrom[*from] = from

	for {
		temp, err := frontier.Remove()
		if err != nil {
			break
		}

		coord := temp.(*weightedCoord)

		if (*coord.coord) == (*to) {
			pathPoint := coord.coord
			for {
				path = append(path, pathPoint)
				if (*pathPoint) == (*from) {
					break
				}
				pathPoint = cameFrom[*pathPoint]
			}
			break
		}

		neighbours := m.GetNeighbours(coord.coord)
		for _, neighbour := range neighbours {
			if _, ok := cameFrom[*neighbour]; !ok && m.IsFree(neighbour) {
				weight := int(math.Abs(float64(to.X-neighbour.X)) + math.Abs(float64(to.Y-neighbour.Y)))
				frontier.Insert(&weightedCoord{coord: neighbour, weight: weight})
				cameFrom[*neighbour] = coord.coord
			}
		}
	}

	return path
}
