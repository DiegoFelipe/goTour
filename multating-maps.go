package main

import (
	"fmt"
)

func mainmapmultating() {
	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println(m["answer"])

	v, ok := m["answer"]
	fmt.Println(v, ok)
}