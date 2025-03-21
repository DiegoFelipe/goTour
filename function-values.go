package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

func mainfuncvalues() {
	hypot := func(x,y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3,3))

	fmt.Println(compute(hypot))
}