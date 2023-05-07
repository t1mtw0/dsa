package main

type MinHeap struct {
	data   []int
	length int
}

func (h *MinHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(h.length)
	h.length++
}

func (h *MinHeap) Delete() int {
	if h.length == 0 {
		return -1
	}

	out := h.data[0]
	h.length--

	if h.length == 0 {
		h.data = []int{}
		return out
	}

	h.data[0] = h.data[h.length]
	h.heapifyDown(0)
	return out
}

func (h *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := h.parent(idx)
	parentV := h.data[p]
	v := h.data[idx]

	if parentV > v {
		h.data[idx] = parentV
		h.data[p] = v
		h.heapifyUp(p)
	}
}

func (h *MinHeap) heapifyDown(idx int) {
	lIdx := h.leftChild(idx)
	rIdx := h.rightChild(idx)

	if idx >= h.length || lIdx >= h.length {
		return
	}

	lV := h.data[lIdx]
	rV := h.data[rIdx]
	v := h.data[idx]

	if lV > rV && v > rV {
		h.data[rIdx] = v
		h.data[idx] = rV
		h.heapifyDown(rIdx)
	} else if rV > lV && v > lV {
		h.data[lIdx] = v
		h.data[idx] = lV
		h.heapifyDown(lIdx)
	}

}

func (h MinHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h MinHeap) leftChild(idx int) int {
	return (idx * 2) + 1
}

func (h MinHeap) rightChild(idx int) int {
	return (idx * 2) + 2
}
