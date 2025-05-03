// n! means n × (n − 1) × ... × 3 × 2 × 1

// For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

// Find the sum of the digits in the number 100!
// Answer:  648
package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	num := big.NewInt(100)
	factorial := factorial(num)
	str_num := factorial.String()
	result := 0
	for _, digit := range str_num {
		digit_int, err := strconv.Atoi(string(digit))
		if err != nil {
			fmt.Println("Error al convertir el digito:", err)
			continue
		}
		result += digit_int
	}
	fmt.Println(str_num)
	fmt.Println(result)
	fmt.Println("Computational time: ", time.Since(start))
}

// factorial  Función Factorial recursiva
func factorial(n *big.Int) *big.Int {
	// Caso base: si n == 1, retorna 1
	if n.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}

	// Crear una copia de n para evitar modificar el valor original
	nMinusOne := new(big.Int).Sub(n, big.NewInt(1))

	// Calcular factorial recursivamente
	return new(big.Int).Mul(n, factorial(nMinusOne))
}
