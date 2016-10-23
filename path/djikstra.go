package path

import "github.com/charithe/algorithms/ds"

func Djikstra(m *Map, from *Coordinate, to *Coordinate) Path {
	var path []*Coordinate
	cameFrom := make(map[Coordinate]*Coordinate)
	frontier := ds.NewCircularQueue(m.size * m.size)
	costs := make(map[Coordinate]int)

	frontier.Enqueue(from)
	cameFrom[*from] = from
	costs[*from] = 0

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
		costSoFar := costs[*coord]
		for _, neighbour := range neighbours {
			// Need Priority queue
			if _, ok := cameFrom[*neighbour]; !ok && m.IsFree(neighbour) {
				frontier.Enqueue(neighbour)
				cameFrom[*neighbour] = coord
			}
		}
	}

	return path

}
