// Em Go, uma interface define um conjunto de métodos que um tipo deve implementar. Elas permitem polimorfismo (a capacidade de tratar diferentes tipos de forma uniforme), sem a necessidade de herança.
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func maininterface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)

	a = f

	fmt.Println(a.Abs())

}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex2 struct {
    X, Y float64
}

func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}