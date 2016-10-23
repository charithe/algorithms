package main

import "github.com/charithe/algorithms/ds"

func main() {
	//m := path.NewMap(10, path.NewObstacle(2, 2, 8, 2), path.NewObstacle(0, 8, 8, 8), path.NewObstacle(8, 2, 8, 8))
	//p := path.BreadthFirst(m, path.NewCoord(0, 0), path.NewCoord(9, 9))
	//m.DrawPath(p)
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
}
