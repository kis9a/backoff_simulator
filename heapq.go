package main

type Heapq []*OccRequest

func (h Heapq) Len() int { return len(h) }

func (h Heapq) Less(i, j int) bool {
	return h[i].time < h[j].time
}

func (h Heapq) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heapq) Push(x interface{}) {
	item := x.(*OccRequest)
	*h = append(*h, item)
}

func (h *Heapq) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func newHeapq() *Heapq {
	return &Heapq{}
}
