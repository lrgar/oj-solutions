package main

import (
	"fmt"
	"sort"
)

type sortInterface []int64

func (p sortInterface) Len() int           { return len(p) }
func (p sortInterface) Less(i, j int) bool { return p[i] > p[j] }
func (p sortInterface) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func solve(time int64, r, c int, b int64, mi, si, pi []int64) bool {
	fill := make(sortInterface, c)
	for i := 0; i < c; i++ {
		if fill[i] = (time - pi[i]) / si[i]; fill[i] > mi[i] {
			fill[i] = mi[i]
		} else if fill[i] < 0 {
			fill[i] = 0
		}
	}

	sort.Sort(fill)

	var cnt int64
	for i := 0; i < r; i++ {
		cnt += fill[i]
	}

	return cnt >= b
}

func main() {
	var t int
	fmt.Scan(&t)

	for ti := 0; ti < t; ti++ {
		var r, c int
		var b int64
		fmt.Scan(&r, &b, &c)

		mi := make([]int64, c)
		si := make([]int64, c)
		pi := make([]int64, c)

		var low, high int64

		for i := 0; i < c; i++ {
			fmt.Scan(&mi[i], &si[i], &pi[i])
			if x := mi[i]*si[i] + pi[i]; x > high {
				high = x
			}
		}

		for low < high {
			if mid := (low + high) / 2; solve(mid, r, c, b, mi, si, pi) {
				high = mid
			} else {
				low = mid + 1
			}
		}

		fmt.Printf("Case #%d: %d\n", ti+1, low)
	}
}
