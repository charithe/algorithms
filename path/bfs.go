package path

import "github.com/charithe/algorithms/ds"

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
