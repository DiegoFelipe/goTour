// Function closures
// Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

// For example, the adder function returns a closure. Each closure is bound to its own sum variable.
// Em Go, closures são funções que capturam e usam variáveis definidas fora de seu escopo imediato. Elas são úteis para encapsular estados e comportamentos de forma elegante.

// 🔹 Como funciona um closure em Go?
// Um closure ocorre quando uma função anônima captura e usa variáveis de fora do seu corpo. Essas variáveis não são copiadas; em vez disso, a função mantém uma referência a elas.
package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
        count++
        return count
    }
}

func mainfclsure() {
	current := counter()

	fmt.Println(current())
	fmt.Println(current())
	fmt.Println(current())

	newCounter := counter()
	fmt.Println(newCounter())
}