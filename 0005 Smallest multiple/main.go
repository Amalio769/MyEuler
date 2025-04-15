// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
// Answer:  232792560
// Algorithm: 1*2⁴*3²*5*7*11*13*17*19
// 1. Find the prime factors of each number from 1 to 20
// 2. Find the maximum exponent of each prime factor
// 3. Multiply the prime factors raised to their maximum exponents
package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	// 1. Find the prime factors of each number from 1 to 20
	n := 20
	factors := make(map[int]int)
	for i := 2; i <= n; i++ {
		num := i
		for j := 2; j <= num; j++ {
			if num%j == 0 {
				factors[j]++
				num /= j
				j--
			}
		}
	}
	// 2. Find the maximum exponent of each prime factor
	maxFactors := make(map[int]int)
	for k, v := range factors {
		if maxFactors[k] < v {
			maxFactors[k] = v
		}
	}
	// 3. Multiply the prime factors raised to their maximum exponents
	result := big.NewInt(1)
	for k, v := range maxFactors {
		result.Mul(result, big.NewInt(int64(math.Pow(float64(k), float64(v)))))
	}
	fmt.Println("Answer: ", result)
	fmt.Println("Execution time: ", time.Since(start))
}
