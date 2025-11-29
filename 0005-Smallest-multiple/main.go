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
	fmt.Println("Factors: ", factors)
	for i := 2; i <= n; i++ {
		if IsPrimeNumber(i) {
			factors[i] = 1
			continue
		}
	}
	fmt.Println("Factors: ", factors)
	// 2. Find the maximum exponent of each prime factor
	for num := 2; num < n; num++ {
		for k, v := range factors {
			if num%k == 0 {
				max_exp := MaxExponentDiv(num, k)
				if max_exp > v {
					factors[k] = max_exp
				}
			}
		}
	}
	fmt.Println("Factors updated with maximum exponent: ", factors)
	// 3. Multiply the prime factors raised to their maximum exponents
	result := big.NewInt(1)
	for k, v := range factors {
		result.Mul(result, big.NewInt(int64(math.Pow(float64(k), float64(v)))))
	}
	fmt.Println("Answer: ", result)
	fmt.Println("Execution time: ", time.Since(start))
}

func IsPrimeNumber(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func MaxExponentDiv(num int, div int) int {
	if num%div == 0 {
		max_exp := 1
		for i := 2; int(math.Pow(float64(div), float64(i))) <= num; i++ {
			if num%int(math.Pow(float64(div), float64(i))) == 0 {
				max_exp++
			}
		}
		return max_exp
	}
	return 0
}
