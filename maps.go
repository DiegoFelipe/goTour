// package main

// import "fmt"

// type Vertex struct {
// 	Lat, Long float64
// }

// var m map[string]Vertex

//	func main() {
//		m = make(map[string]Vertex)
//		m["Bell Labs"] = Vertex{
//			40.68433, -74.39967,
//		}
//		fmt.Println(m["Bell Labs"])
//	}
package main

import "fmt"

type Position struct {
	Lat, Long float64
}

func mainmaps() {
	m := make(map[string]Position)
	m["Bell Labs"] = Position{40.654546, -74.312423}

	fmt.Println(m["Bell Labs"])
}