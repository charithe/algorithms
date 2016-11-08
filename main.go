package main

import (
	"fmt"

	"github.com/charithe/algorithms/path"
)

func main() {
	pathFinders := []struct {
		name string
		pf   path.PathFinder
	}{
		{"Breadth-first", path.BreadthFirst},
		{"Greedy breadth-first", path.GreedyBreadthFirst},
		{"Djikstra", path.Djikstra},
		{"A*", path.AStar},
	}

	m := path.NewMap(10, path.NewObstacle(2, 2, 8, 2), path.NewObstacle(1, 8, 8, 8), path.NewObstacle(8, 2, 8, 8))
	from := path.NewCoord(4, 7)
	to := path.NewCoord(9, 5)

	for _, pf := range pathFinders {
		fmt.Println(pf.name)
		p := pf.pf(m, from, to)
		m.DrawPath(p)
	}
}
