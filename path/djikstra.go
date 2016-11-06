package path

import "github.com/charithe/algorithms/ds"

func Djikstra(m *Map, from *Coordinate, to *Coordinate) Path {
	var path []*Coordinate
	cameFrom := make(map[Coordinate]*Coordinate)
	frontier := ds.NewBinaryHeap(weightedCoordComparer, m.size*m.size)
	costs := make(map[Coordinate]int)

	frontier.Insert(&weightedCoord{coord: from, weight: 0})
	cameFrom[*from] = from
	costs[*from] = 0

	for {
		temp, err := frontier.Remove()
		if err != nil {
			break
		}

		wc := temp.(*weightedCoord)

		if (*wc.coord) == (*to) {
			pathPoint := wc.coord
			for {
				path = append(path, pathPoint)
				if (*pathPoint) == (*from) {
					break
				}
				pathPoint = cameFrom[*pathPoint]
			}
			break
		}

		neighbours := m.GetNeighbours(wc.coord)
		for _, neighbour := range neighbours {
			if m.IsFree(neighbour) {
				newCost := costs[*wc.coord] + 1
				if c, ok := costs[*neighbour]; !ok || newCost < c {
					costs[*neighbour] = newCost
					frontier.Insert(&weightedCoord{coord: neighbour, weight: newCost})
					cameFrom[*neighbour] = wc.coord
				}
			}
		}
	}
	return path

}
