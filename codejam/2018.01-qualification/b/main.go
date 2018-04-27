package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var stdin = bufio.NewReader(os.Stdin)

func main() {
	var t int
	fmt.Fscan(stdin, &t)

	for ti := 0; ti < t; ti++ {
		var n int
		fmt.Fscan(stdin, &n)

		all := make([]int, n)
		even := make([]int, (n+1)/2)
		odd := make([]int, n/2)

		for i, j, k := 0, 0, 0; i < n; i++ {
			fmt.Fscan(stdin, &all[i])
			if i%2 == 0 {
				even[j] = all[i]
				j++
			} else {
				odd[k] = all[i]
				k++
			}
		}

		sort.Ints(all)
		sort.Ints(even)
		sort.Ints(odd)

		i := 0
		for j, k := 0, 0; i < n; i++ {
			if i%2 == 0 {
				if even[j] != all[i] {
					break
				}
				j++
			} else {
				if odd[k] != all[i] {
					break
				}
				k++
			}
		}

		if i == n {
			fmt.Printf("Case #%d: OK\n", ti+1)
		} else {
			fmt.Printf("Case #%d: %d\n", ti+1, i)
		}
	}
}
