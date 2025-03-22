// package main

// import "fmt"

// type Vertex struct {
// 	Lat, Long float64
// }

// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -74.39967,
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,
// 	},
// }

// func main() {
// 	fmt.Println(m)
// }

package main

import "fmt"

type Position struct {
	Lat, Long float64
}

var m = map[string]Position{
	"Bells labs": Position{
		40.654546, -74.312423,
	},
	"Google": Position{
        37.42202, -122.08408,
    },
}

func mainmmapsliteral() {
	fmt.Println(m)
}