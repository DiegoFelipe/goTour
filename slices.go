// Slices
// An array has a fixed size.
// A slice, on the other hand, is a dynamically-sized, flexible view
// into the elements of an array.
// In practice, slices are much more common than arrays.

package main

import (
	"fmt"
)

func main2343() {
	primes := [6]int{2,3,5,7,9,11}

	slice := primes[1:4]
	fmt.Println(slice)
}