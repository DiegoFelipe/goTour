// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		t := time.Now()
//		switch {
//		case t.Hour() < 12:
//			fmt.Println("Good morning!")
//		case t.Hour() < 17:
//			fmt.Println("Good afternoon.")
//		default:
//			fmt.Println("Good evening.")
//		}
//	}
package main

import (
	"fmt"
	"time"
)

func main4() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
