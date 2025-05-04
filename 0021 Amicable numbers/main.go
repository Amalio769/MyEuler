// Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
// If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

// For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The
// proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

// Evaluate the sum of all the amicable numbers under 10000.
// Answer:  31626
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(AmicableNumbers(220))
	fmt.Println("Computational time: ", time.Since(start))
}

func ProperDivisors(num int) []int {
	divisors := []int{}
	for i := 1; i < num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func SumSlice(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}

func AmicableNumbers(num int) []int {
	result := []int{}
	num1 := SumSlice(ProperDivisors(num))
	if SumSlice(ProperDivisors(num1)) == num {
		result = append(result, num, num1)
	}
	return result
}
