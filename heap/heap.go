package heap

type Heap struct {
	heap []int
}

func (h *Heap) Left(i int) int {
	return i * 2
}

func (h *Heap) Right(i int) int {
	return (i * 2) + 1
}

func (h *Heap) Parent(i int) int {
	return i / 2
}

func (h *Heap) Set(i, x int) {
	h.heap[i] = x
}

func (h *Heap) Get(i int) int {
	return h.heap[i]
}

func (h *Heap) Size() int {
	return len(h.heap)
}

func MaxHeapify(h Heap, i int) {
	largest := 0
	l := h.Left(i)
	r := h.Right(i)
	if l <= h.Size() && h.Get(l) > h.Get(i) {
		largest = l
	} else {
		largest = i
	}
	if r > largest {
		largest = r
	}
	if largest != i {
		k := h.Get(i)
		h.Set(i, h.Get(largest))
		h.Set(largest, k)
		MaxHeapify(h, largest)
	}
}
