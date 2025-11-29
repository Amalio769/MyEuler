// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

// Find the sum of all the primes below two million.
// Answer:  142913828922
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	primes_below := 2000000
	var list_prime []int
	for i := 2; i < primes_below; i++ {
		next_prime := NextPrime(list_prime)
		if next_prime < primes_below {
			list_prime = append(list_prime, next_prime)
		} else {
			break
		}
		fmt.Println(next_prime)
	}
	sum := 0
	for _, num := range list_prime {
		sum += num
	}
	fmt.Println(sum)

	fmt.Println("Computational time: ", time.Since(start))
}

func NextPrime(list_prime []int) int {
	if len(list_prime) == 0 {
		return 2
	}
	for num := list_prime[len(list_prime)-1] + 1; ; num++ {
		num_is_prime := true
		for _, pf := range list_prime {
			if num%pf == 0 {
				num_is_prime = false
				break
			}
		}
		if num_is_prime {
			return num
		}
	}
}
