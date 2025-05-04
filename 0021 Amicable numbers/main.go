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
	list_amicable_numbers := []int{}
	for i := 1; i < 10000; i++ {
		if num1 := AmicableNumbers(i); num1 != 0 {
			if !IsNumInSlice(list_amicable_numbers, i) {
				list_amicable_numbers = append(list_amicable_numbers, i)
			}
			if !IsNumInSlice(list_amicable_numbers, num1) {
				list_amicable_numbers = append(list_amicable_numbers, num1)
			}
		}
	}
	fmt.Println(list_amicable_numbers)
	fmt.Println("La suma de los números amigable es: ", SumSlice(list_amicable_numbers))
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

func AmicableNumbers(num int) int {
	num1 := SumSlice(ProperDivisors(num))
	if SumSlice(ProperDivisors(num1)) == num && num != num1 {
		return num1
	}
	return 0
}

func IsNumInSlice(slice []int, num int) bool {
	for _, value := range slice {
		if value == num {
			return true
		}
	}
	return false
}
