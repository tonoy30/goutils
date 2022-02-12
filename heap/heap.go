package heap

type IntMaxHeap []int

func (h IntMaxHeap) Len() int {
	return len(h)
}
func (h IntMaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h IntMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}
