package main

import (
	"fmt"
	"runtime"
)

func main2 () {
	os := runtime.GOOS
	fmt.Println(os)
}