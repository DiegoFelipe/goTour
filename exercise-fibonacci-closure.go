package main

import "fmt"

func fibonacci() func() int {
	a,b := 0,1
	return func() int {
		result := a
		a,b = b, a+b
		return result
	}
}

func mainexfib() {
	fib := fibonacci()

	for i := 0; i < 12; i++ {
		fmt.Println(fib())
	}
}