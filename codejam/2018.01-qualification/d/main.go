package main

import (
	"fmt"
	"math"
)

// WA for hidden test.

var pi4 = math.Pi / 4

func rt(r1, r2 float64, p [3]float64) [3]float64 {
	return rotPoint(rotX(r2), rotPoint(rotY(r1), p))
}

func rotX(r float64) [3][3]float64 {
	return [3][3]float64{
		{1, 0, 0},
		{0, math.Cos(r), -math.Sin(r)},
		{0, math.Sin(r), math.Cos(r)},
	}
}

func rotY(r float64) [3][3]float64 {
	return [3][3]float64{
		{math.Cos(r), 0, math.Sin(r)},
		{0, 1, 0},
		{-math.Sin(r), 0, math.Cos(r)},
	}
}

func rotPoint(m [3][3]float64, p [3]float64) [3]float64 {
	return [3]float64{
		m[0][0]*p[0] + m[0][1]*p[1] + m[0][2]*p[2],
		m[1][0]*p[0] + m[1][1]*p[1] + m[1][2]*p[2],
		m[2][0]*p[0] + m[2][1]*p[1] + m[2][2]*p[2],
	}
}

func areaParallelogram(a, b, c [3]float64) float64 {
	return math.Abs(a[0]*b[1] + b[0]*c[1] + c[0]*a[1] - a[1]*b[0] - b[1]*c[0] - c[1]*a[0])
}

func calculateArea(r1, r2 float64) float64 {
	a := rt(r1, r2, [3]float64{0.5, -0.5, 0.5})
	b := rt(r1, r2, [3]float64{0.5, -0.5, -0.5})
	c := rt(r1, r2, [3]float64{0.5, 0.5, -0.5})
	d := rt(r1, r2, [3]float64{-0.5, -0.5, -0.5})

	area := areaParallelogram(a, b, c) +
		areaParallelogram(c, b, d) +
		areaParallelogram(a, b, d)

	return area
}

func calculatePoints(r1, r2 float64) [3][3]float64 {
	a := rt(r1, r2, [3]float64{0.5, 0, 0})
	b := rt(r1, r2, [3]float64{0, 0, 0.5})
	c := rt(r1, r2, [3]float64{0, 0.5, 0})
	return [3][3]float64{a, b, c}
}

func main() {
	var t int
	fmt.Scan(&t)

	for ti := 0; ti < t; ti++ {
		var a float64
		fmt.Scan(&a)

		var r1, r2 float64
		if a <= math.Sqrt2 {
			r2 = 0
			low, high := 0.0, pi4
			for high-low > 1e-12 {
				mid := (low + high) / 2
				if calculateArea(mid, r2) < a {
					low = mid
				} else {
					high = mid
				}
			}
			r1 = low
		} else {
			r1 = pi4
			low, high := 0.0, pi4
			for high-low > 1e-12 {
				mid := (low + high) / 2
				if calculateArea(r1, mid) < a {
					low = mid
				} else {
					high = mid
				}
			}
			r2 = low
		}

		pts := calculatePoints(r1, r2)

		fmt.Printf("Case #%d:\n", ti+1)
		for i := 0; i < 3; i++ {
			fmt.Printf("%.16f %.16f %.16f\n", pts[i][0], pts[i][2], pts[i][1])
		}
	}
}
