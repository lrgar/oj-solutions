package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for ti := 1; ti <= t; ti++ {
		var d float64
		var n int
		fmt.Scan(&d, &n)

		max := 0.0
		for i := 0; i < n; i++ {
			var k, s float64
			fmt.Scan(&k, &s)

			if time := (d - k) / s; time > max {
				max = time
			}
		}

		fmt.Printf("Case #%d: %.6f\n", ti, d/max)
	}
}
