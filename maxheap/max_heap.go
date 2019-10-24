package heap

type MaxHeap struct {
	heap []int
}

func (h *MaxHeap) Left(i int) int {
	return (i * 2) + 1
}

func (h *MaxHeap) Right(i int) int {
	return (i * 2) + 2
}

func (h *MaxHeap) Parent(i int) int {
	return i / 2
}

func (h *MaxHeap) Set(i, x int) {
	h.heap[i] = x
}

func (h *MaxHeap) Get(i int) int {
	return h.heap[i]
}

func (h *MaxHeap) Size() int {
	return len(h.heap) - 1
}

func (h *MaxHeap) Array() []int {
	return h.heap
}

func MaxHeapify(h MaxHeap, i int) {
	largest := 0
	l := h.Left(i)
	r := h.Right(i)
	if l <= h.Size() && h.Get(l) > h.Get(i) {
		largest = l
	} else {
		largest = i
	}
	if r <= h.Size() && h.Get(r) > h.Get(largest) {
		largest = r
	}
	if largest != i {
		k := h.Get(i)
		h.Set(i, h.Get(largest))
		h.Set(largest, k)
		MaxHeapify(h, largest)
	}
}

// BuildMaxHeap builds a max heap from a given array.
func BuildMaxHeap(input []int) *MaxHeap {
	h := MaxHeap{
		heap: input,
	}
	for i := (len(input) - 1) / 2; i >= 0; i-- {
		MaxHeapify(h, i)
	}
	return &h
}

// Sort performs a heap-sort of array input.
func Sort(input []int) []int {
	heap := BuildMaxHeap(input)
	result := make([]int, len(input))
	for i := heap.Size(); i > 0; i-- {
		result[i] = heap.Get(0)
		heap.Set(0, heap.Get(i))
		heap.heap = heap.heap[:i]
		MaxHeapify(*heap, 0)
	}
	result[0] = heap.Get(0)
	return result
}
