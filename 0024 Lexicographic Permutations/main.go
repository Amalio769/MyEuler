// A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1,
//  2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order.
//  The lexicographic permutations of 0, 1 and 2 are:

// 012   021   102   120   201   210

// What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
// Answer:  2783915460
package main

import (
	"fmt"
	"time"
)

// Función para calcular el factorial de un número
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Función para encontrar la permutación lexicográfica en una iteración específica
func findNthPermutation(digits []int, n int) string {
	result := ""
	k := n - 1 // Convertir a índice basado en 0

	for len(digits) > 0 {
		// Calcular el factorial del tamaño actual del conjunto - 1
		fact := factorial(len(digits) - 1)

		// Determinar el índice del dígito actual
		index := k / fact
		result += fmt.Sprintf("%d", digits[index])

		// Eliminar el dígito usado
		digits = append(digits[:index], digits[index+1:]...)

		// Actualizar k para la siguiente iteración
		k %= fact
	}

	return result
}

func main() {
	start := time.Now()

	// Dígitos disponibles
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Iteración deseada (por ejemplo, la décima)
	iteration := 1000000
	permutation := findNthPermutation(digits, iteration)

	fmt.Printf("La permutación en la iteración %d es: %s\n", iteration, permutation)
	fmt.Println("Tiempo de cómputo:", time.Since(start))
}
