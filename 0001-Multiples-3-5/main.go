// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these
// multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.
// Answer: 233168
// Completed on Thu, 05 May 2024, 08:25

package main

import "fmt"

func Multiples_3_5(value int) bool {
	result := false
	if (value%3 == 0) || (value%5 == 0) {
		result = true
	}
	return result
}

func main() {
	result := 0
	for i := 1; i < 1000; i++ {
		if Multiples_3_5(i) {
			result += i
		}
	}
	fmt.Println(result)
}
