package main333

import "fmt"

type Vertex struct {
	X, Y int
}

func main4s() {
    v1 := Vertex{1,2}
	pv1 := &v1
	pv1.X = 1e9
	fmt.Println(v1)
}