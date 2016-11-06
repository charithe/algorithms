package main

import (
	"fmt"

	"github.com/charithe/algorithms/path"
)

func main() {
	fmt.Println("Breadth-first")
	m := path.NewMap(10, path.NewObstacle(2, 2, 8, 2), path.NewObstacle(1, 8, 8, 8), path.NewObstacle(8, 2, 8, 8))
	pb := path.BreadthFirst(m, path.NewCoord(4, 7), path.NewCoord(9, 5))
	m.DrawPath(pb)
	fmt.Println()
	fmt.Println("Greedy breadth-first")
	pg := path.GreedyBreadthFirst(m, path.NewCoord(4, 7), path.NewCoord(9, 5))
	m.DrawPath(pg)
	fmt.Println()
	fmt.Println("Djikstra")
	pd := path.Djikstra(m, path.NewCoord(4, 7), path.NewCoord(9, 5))
	m.DrawPath(pd)
	/*
		c := func(a interface{}, b interface{}) int {
			ai := a.(int)
			bi := b.(int)
			return ai - bi
		}

		bt := ds.NewBinaryTree(c)
		bt.Insert(10, 20, 30, 40, 50, 60, 70, 80, 90)
		bt.Print()
		bt.Balance()
		bt.Print()
	*/
}
