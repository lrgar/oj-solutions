package main

import (
	"fmt"
)

func solve(r, c int) {
	var visited [20][10]bool
	var effect [20][10]int

	for i := 1; i < r-1; i++ {
		for j := 1; j < c-1; j++ {
			effect[i][j] = 9
		}
	}

	for e := 9; e > 0; e-- {
		for i := 1; i < r-1; i++ {
			for j := 1; j < c-1; j++ {
				if effect[i][j] == e {
					for {
						fmt.Println(i+1, j+1)

						var ip, jp int
						fmt.Scan(&ip, &jp)

						if ip == -1 && jp == -1 {
							panic("FAILED TEST.")
						}

						if ip == 0 && jp == 0 {
							return
						}

						ip--
						jp--

						if visited[ip][jp] {
							continue
						}

						visited[ip][jp] = true

						for di := -1; di <= 1; di++ {
							for dj := -1; dj <= 1; dj++ {
								in, jn := ip+di, jp+dj
								if in < 1 || jn < 1 || in >= r-1 || jn >= c-1 {
									continue
								}
								effect[in][jn]--
							}
						}

						break
					}
				}
			}
		}
	}

	panic("FAILED TO SOLVE.")
}

func main() {
	var t int
	fmt.Scan(&t)

	for ti := 0; ti < t; ti++ {
		var a int
		fmt.Scan(&a)

		var w, h int
		if a == 10 {
			w, h = 4, 3
		} else if a == 20 {
			w, h = 5, 4
		} else if a == 200 {
			w, h = 20, 10
		}

		solve(w, h)
	}
}
