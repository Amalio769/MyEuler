// The Fibonacci sequence is defined by the recurrence relation:
// Fn=Fn-1 + Fn-2, where F1=1 and F2=1
// Hence the first 12 terms will be:
// F1 = 1
// F2 = 1
// F3 = 2
// F4 = 3
// F5 = 5
// F6 = 8
// F7 = 13
// F8 = 21
// F9 = 34
// F10 = 55
// F11 = 89
// F12 = 144
// The 12th term, F12, is the first term to contain three digits.
// What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
// Answer:4782
package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()

	// Encontrar el índice del primer término de Fibonacci con 1000 dígitos
	digits := 1000
	index := FibonacciIndexWithDigits(digits)

	fmt.Printf("El índice del primer término de Fibonacci con %d dígitos es: %d\n", digits, index)
	fmt.Println("Tiempo de cómputo:", time.Since(start))
}

// Función para encontrar el índice del primer término de Fibonacci con al menos 'digits' dígitos
func FibonacciIndexWithDigits(digits int) int {
	// Inicializar los primeros términos de Fibonacci
	a := big.NewInt(1) // F1
	b := big.NewInt(1) // F2
	index := 2         // Comenzamos en el segundo término

	// Crear un límite de comparación: 10^(digits-1)
	limit := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(digits-1)), nil)

	// Iterar hasta que el número de Fibonacci tenga al menos 'digits' dígitos
	for b.Cmp(limit) < 0 {
		// Calcular el siguiente término de Fibonacci
		a.Add(a, b) // a = a + b
		a, b = b, a // Intercambiar a y b
		index++
	}

	return index
}
