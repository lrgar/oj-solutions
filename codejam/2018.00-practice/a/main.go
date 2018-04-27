package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for ti := 0; ti < t; ti++ {
		var a, b, n int64
		fmt.Scan(&a, &b, &n)

		a++
		for {
			c := (a + b) / 2
			fmt.Println(c)

			var res string
			fmt.Scan(&res)

			if res == "CORRECT" {
				break
			} else if res == "TOO_SMALL" {
				a = c + 1
			} else if res == "TOO_BIG" {
				b = c - 1
			}
		}
	}
}
