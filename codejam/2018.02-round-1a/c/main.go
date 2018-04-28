package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scan(&t)

	for ti := 0; ti < t; ti++ {
		var n, p int
		fmt.Scan(&n, &p)

		m := n * 250
		dp := make([]float64, m+1)
		for i := 1; i <= m; i++ {
			dp[i] = -1
		}

		var z int

		for i := 0; i < n; i++ {
			var wi, hi int
			fmt.Scan(&wi, &hi)
			if hi < wi {
				wi, hi = hi, wi
			}

			h := math.Sqrt(float64(wi*wi+hi*hi)) - float64(wi)
			z += 2 * (wi + hi)

			for j := m; j >= 0; j-- {
				if dp[j] != -1 {
					dp[j+wi] = math.Max(dp[j+wi], dp[j]+h)
				}
			}
		}

		pr := p - z

		var out float64
		for j := 0; j <= m && 2*j <= pr; j++ {
			if dp[j] == -1 {
				continue
			}

			var x float64
			if r := 2 * (dp[j] + float64(j)); r < float64(pr) {
				x = r
			} else {
				x = float64(pr)
			}

			if x > out {
				out = x
			}
		}

		fmt.Printf("Case #%d: %.8f\n", ti+1, out+float64(z))
	}
}
