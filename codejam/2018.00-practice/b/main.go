package main

import (
	"fmt"
	"sort"
)

type SortInterface struct {
	pos, cnt []int
}

func (p SortInterface) Len() int           { return len(p.pos) }
func (p SortInterface) Less(i, j int) bool { return p.cnt[p.pos[i]] > p.cnt[p.pos[j]] }
func (p SortInterface) Swap(i, j int)      { p.pos[i], p.pos[j] = p.pos[j], p.pos[i] }

func main() {
	ids := make([]string, 26)
	for i := 0; i < 26; i++ {
		ids[i] = string([]byte{byte('A') + byte(i)})
	}

	var t int
	fmt.Scan(&t)

	for ti := 0; ti < t; ti++ {
		var n int
		fmt.Scan(&n)

		sum := 0
		cnt := make([]int, n)
		pos := make([]int, n)
		for i := 0; i < n; i++ {
			pos[i] = i
			fmt.Scan(&cnt[i])
			sum += cnt[i]
		}

		fmt.Printf("Case #%d:", ti+1)

		for sum > 0 {
			sort.Sort(SortInterface{pos, cnt})

			a, b := pos[0], pos[1]
			if sum%2 == 1 {
				fmt.Printf(" %s", ids[a])
				cnt[a]--
				sum--
			} else {
				fmt.Printf(" %s%s", ids[a], ids[b])
				cnt[a]--
				cnt[b]--
				sum -= 2
			}
		}

		fmt.Println()
	}
}
