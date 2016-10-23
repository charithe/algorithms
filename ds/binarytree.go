package ds

import (
	"fmt"
	"math"
)

type Comparer func(interface{}, interface{}) int

type treeNode struct {
	data  interface{}
	left  *treeNode
	right *treeNode
}

func insert(root *treeNode, data interface{}, compare Comparer) *treeNode {
	if root == nil {
		return &treeNode{data: data}
	}

	if compare(data, root.data) > 0 {
		root.right = insert(root.right, data, compare)
	} else {
		root.left = insert(root.left, data, compare)
	}

	return root
}

func swapData(a *treeNode, b *treeNode) {
	tmp := a.data
	a.data = b.data
	b.data = tmp
}

func dswTreeToVine(root *treeNode) {
	tail := root
	rest := tail.right

	for rest != nil {
		if rest.left == nil {
			tail = rest
			rest = rest.right
		} else {
			temp := rest.left
			rest.left = temp.right
			temp.right = rest
			rest = temp
			tail.right = temp
		}
	}
}

func dswVineToTree(root *treeNode, size int) {
	leaves := size + 1 - int(math.Pow(2, math.Floor(math.Log2(float64(size+1)))))
	dswCompress(root, leaves)
	size = size - leaves
	for size > 1 {
		size = size / 2
		dswCompress(root, size)
	}
}

func dswCompress(root *treeNode, count int) {
	scanner := root
	for i := 0; i < count; i++ {
		child := scanner.right
		scanner.right = child.right
		scanner = scanner.right
		child.right = scanner.left
		scanner.left = child
	}
}

func printTree(root *treeNode, indent string, rightChild bool) {
	fmt.Print(indent)
	var newIndent string
	if rightChild {
		fmt.Print(" └╴ ")
		newIndent = indent + "   "
	} else {
		fmt.Print(" ├╴ ")
		newIndent = indent + " │ "
	}
	if root == nil {
		fmt.Print("<>\n")
		return
	} else {
		fmt.Printf("%v\n", root.data)
		printTree(root.left, newIndent, false)
		printTree(root.right, newIndent, true)
	}
}

func height(root *treeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.left)
	rightHeight := height(root.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func toArray(root *treeNode) []interface{} {
	treeHeight := height(root)
	arraySize := (treeHeight * 2) + 1
	array := make([]interface{}, arraySize)
	addToArray(array, root, 0)
	return array
}

func addToArray(array []interface{}, node *treeNode, level int) {
	if node == nil {
		return
	}

	index := level * 2
	array[index] = node
	array[level+1] = node.left
	array[level+2] = node.right
	addToArray(array, node.left, level+1)
	addToArray(array, node.right, level+1)
}

type BinaryTree struct {
	root    *treeNode
	size    int
	compare Comparer
}

func NewBinaryTree(comparer Comparer) *BinaryTree {
	return &BinaryTree{compare: comparer}
}

func (bt *BinaryTree) Insert(data ...interface{}) {
	for _, d := range data {
		bt.root = insert(bt.root, d, bt.compare)
		bt.size += 1
	}
}

func (bt *BinaryTree) Print() {
	printTree(bt.root, "", true)
}

func (bt *BinaryTree) Balance() {
	pseudoRoot := &treeNode{}
	pseudoRoot.right = bt.root
	dswTreeToVine(pseudoRoot)
	dswVineToTree(pseudoRoot, bt.size)
	bt.root = pseudoRoot.right
}
