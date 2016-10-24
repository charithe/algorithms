package ds

import "fmt"

var (
	ErrHeapFull  = fmt.Errorf("Heap is full")
	ErrHeapEmpty = fmt.Errorf("Heap is empty")
)

type BinaryHeap struct {
	heap           []interface{}
	comparer       Comparer
	size           int
	insertionPoint int
}

func NewBinaryHeap(comparer Comparer, size int) *BinaryHeap {
	return &BinaryHeap{
		heap:           make([]interface{}, size),
		comparer:       comparer,
		size:           size,
		insertionPoint: 0,
	}
}

func (bh *BinaryHeap) Size() int {
	return bh.insertionPoint
}

func (bh *BinaryHeap) Peek() interface{} {
	return bh.heap[0]
}

func (bh *BinaryHeap) Insert(data ...interface{}) error {
	if bh.insertionPoint+len(data) > bh.size {
		return ErrHeapFull
	}

	for _, d := range data {
		bh.heap[bh.insertionPoint] = d
		bh.percolateUp()
		bh.insertionPoint += 1
	}

	return nil
}

func (bh *BinaryHeap) percolateUp() {
	currPtr := bh.insertionPoint
	for {
		parent := (currPtr - 1) / 2
		if bh.comparer(bh.heap[parent], bh.heap[currPtr]) >= 0 {
			break
		}
		bh.swapItems(parent, currPtr)
		currPtr = parent
	}
}

func (bh *BinaryHeap) Remove() (interface{}, error) {
	if bh.insertionPoint == 0 {
		return nil, ErrHeapEmpty
	}

	item := bh.heap[0]
	bh.insertionPoint = bh.insertionPoint - 1
	bh.heap[0] = bh.heap[bh.insertionPoint]
	if bh.insertionPoint > 1 {
		bh.percolateDown()
	}
	return item, nil
}

func (bh *BinaryHeap) percolateDown() {
	currPtr := 0
	for {
		child := (currPtr * 2) + 1
		if child >= bh.insertionPoint || bh.comparer(bh.heap[currPtr], bh.heap[child]) >= 0 {
			break
		}
		bh.swapItems(currPtr, child)
		currPtr = child
	}
}

func (bh *BinaryHeap) swapItems(a int, b int) {
	tmp := bh.heap[a]
	bh.heap[a] = bh.heap[b]
	bh.heap[b] = tmp
}

func (bh *BinaryHeap) String() string {
	return fmt.Sprintf("{ip: %d, heap: %#v}", bh.insertionPoint, bh.heap)
}
