// A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the
// sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

// A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum
// exceeds n.

// As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two
// abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as
// the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is
// known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

// Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
// Answer:  4179871
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	list_abundant_number := []int{}
	for num := 12; num < 28123; num++ {
		if IsAbundantNumber(num) {
			list_abundant_number = append(list_abundant_number, num)
		}
	}
	list_sum_two_abundant_number := []int{}
	for idx, value := range list_abundant_number {
		for _, num := range list_abundant_number[idx:] {
			if value+num < 28123 {
				list_sum_two_abundant_number = append(list_sum_two_abundant_number, value+num)
			}

		}

	}
	list_number_until_28123 := []int{}
	for value := range 28123 {
		list_number_until_28123 = append(list_number_until_28123, value)
	}
	list_number_not_sum_two_abundant_number := []int{}
	for _, value := range list_number_until_28123 {
		if !IsNumInSlice(value, list_sum_two_abundant_number) {
			list_number_not_sum_two_abundant_number = append(list_number_not_sum_two_abundant_number, value)
		}
	}
	// fmt.Println(list_abundant_number)
	// fmt.Println(list_number_not_sum_two_abundant_number)
	fmt.Println(SumSlice(list_number_not_sum_two_abundant_number))
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

func SumSlice(list_numbers []int) int {
	result := 0
	for _, num := range list_numbers {
		result += num
	}
	return result
}

func IsAbundantNumber(num int) bool {
	return SumSlice(ProperDivisors(num)) > num
}

func IsNumInSlice(num int, slice []int) bool {
	for _, value := range slice {
		if num == value {
			return true
		}
	}
	return false
}
