package main

import (
	"fmt"
	"time"
)

func main3() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today. ", time.Saturday + 1)
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.", today, today + 1)
	}
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("When's Saturday?")
// 	today := time.Now().Weekday()
// 	fmt.Println(today + 1, time.Saturday)
// }

