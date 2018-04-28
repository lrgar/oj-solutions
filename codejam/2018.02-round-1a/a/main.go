package main

import (
	"fmt"
)

func solve(r, c, h, v int, grid [][]int) bool {
	if grid[r][c]%((h+1)*(v+1)) != 0 {
		return false
	}

	cc := grid[r][c] / (v + 1)
	rc := grid[r][c] / (h + 1)
	ac := grid[r][c] / (v + 1) / (h + 1)

	vl := make([]int, v+2)
	hl := make([]int, h+2)
	vl[v+1], hl[h+1] = c, r

	for j := 1; j <= v; j++ {
		vl[j] = vl[j-1] + 1

		if vl[j] > c {
			return false
		}

		for ; vl[j] <= c; vl[j]++ {
			if x := grid[r][vl[j]] - grid[r][vl[j-1]]; x > cc {
				return false
			} else if x == cc {
				break
			}
		}
	}

	for j := 1; j <= h; j++ {
		hl[j] = hl[j-1] + 1

		if hl[j] > r {
			return false
		}

		for ; hl[j] <= r; hl[j]++ {
			if x := grid[hl[j]][c] - grid[hl[j-1]][c]; x > rc {
				return false
			} else if x == rc {
				break
			}
		}
	}

	for i := 1; i <= h+1; i++ {
		for j := 1; j <= v+1; j++ {
			if x := grid[hl[i]][vl[j]] - grid[hl[i-1]][vl[j]] - grid[hl[i]][vl[j-1]] + grid[hl[i-1]][vl[j-1]]; x != ac {
				return false
			}
		}
	}

	return true
}

func main() {
	var t int
	fmt.Scan(&t)

	for ti := 0; ti < t; ti++ {
		var r, c, h, v int
		fmt.Scan(&r, &c, &h, &v)

		grid := make([]string, r)
		for i := range grid {
			fmt.Scan(&grid[i])
		}

		cnt := make([][]int, r+1)
		cnt[0] = make([]int, c+1)
		for i := 1; i <= r; i++ {
			cnt[i] = make([]int, c+1)
			for j := 1; j <= c; j++ {
				cnt[i][j] = cnt[i][j-1] + cnt[i-1][j] - cnt[i-1][j-1]
				if grid[i-1][j-1] == '@' {
					cnt[i][j]++
				}
			}
		}

		fmt.Printf("Case #%d: ", ti+1)
		if solve(r, c, h, v, cnt) {
			fmt.Println("POSSIBLE")
		} else {
			fmt.Println("IMPOSSIBLE")
		}
	}
}
