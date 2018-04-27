package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for ti := 0; ti < t; ti++ {
		var d int
		var p string
		fmt.Scan(&d, &p)

		var shots [31]int

		attack := 0
		charged := 0
		shotsCount := 0
		for _, c := range p {
			if c == 'C' {
				charged++
			} else {
				shots[charged]++
				shotsCount++
				attack += 1 << uint(charged)
			}
		}

		if shotsCount > d {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", ti+1)
			continue
		}

		hacks := 0
		for attack > d {
			hacks++
			for i := 30; i > 0; i-- {
				if shots[i] > 0 {
					shots[i]--
					shots[i-1]++
					attack -= 1 << uint(i-1)
					break
				}
			}
		}

		fmt.Printf("Case #%d: %d\n", ti+1, hacks)
	}
}
