// The sum of the squares of the first ten natural numbers is, (1)2+(2)2+...+(10)2 = 385

// The square of the sum of the first ten natural numbers is, (1+2+...+10)2 = (55)2 = 3025

// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is .
// 3025 -385 = 2640
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
// Answer:  25164150
package main

import (
	"fmt"
	"time"
)

func SumOfSquare(num int) int {
	result := 0
	for i := 1; i < num+1; i++ {
		result = result + i*i
	}
	return result
}
func SquareOfSum(num int) int {
	result := 0
	for i := 1; i < num+1; i++ {
		result = result + i
	}
	return result * result
}

func main() {
	start := time.Now()
	fmt.Println("Result :", SquareOfSum(100)-SumOfSquare(100))
	fmt.Println("Tiempo de ejecución: ", time.Since(start))
}
