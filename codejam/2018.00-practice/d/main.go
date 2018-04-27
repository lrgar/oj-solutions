package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int64))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var t int
	fmt.Scan(&t)

	for ti := 1; ti <= t; ti++ {
		var n, k int64
		fmt.Scan(&n, &k)

		h := &IntHeap{n}
		heap.Init(h)

		cnt := map[int64]int64{n: 1}

		var max, min int64
		for k > 0 {
			x := heap.Pop(h).(int64)
			y := cnt[x]
			if a, b := (x-1)/2, x/2; k <= y {
				min, max = a, b
			} else {
				if _, ok := cnt[a]; !ok {
					heap.Push(h, a)
				}
				cnt[a] += y
				if _, ok := cnt[b]; !ok {
					heap.Push(h, b)
				}
				cnt[b] += y
			}
			k -= y
		}

		fmt.Printf("Case #%d: %d %d\n", ti, max, min)
	}
}
