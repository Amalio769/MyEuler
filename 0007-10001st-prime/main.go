// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

// What is the 10 001st prime number?
// Answer:  104743
// Completed on Mon, 15 Aug 2022, 15:54
// Using the function IsPrimeFactor the computational time is around 4.533 sec
// Using the function NextPrimeFactor the computational time is around 0.635 sec
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// number := 2
	// var list_prime []int
	// for {
	// 	if IsPrimeFactor(number) {
	// 		list_prime = append(list_prime, number)
	// 		if len(list_prime) == 10001 {
	// 			break
	// 		}
	// 	}
	// 	number++
	// }
	// fmt.Println(list_prime[len(list_prime)-1])

	var list_prime []int
	for {
		list_prime = append(list_prime, NextPrimeFactor(list_prime))
		if len(list_prime) == 10001 {
			break
		}
	}
	fmt.Println(list_prime[len(list_prime)-1])

	fmt.Println("Tiempo de ejecución: ", time.Since(start))
}

func IsPrimeFactor(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func NextPrimeFactor(list_prime_factor []int) int {
	if len(list_prime_factor) == 0 {
		return 2
	}
	for num := list_prime_factor[len(list_prime_factor)-1] + 1; ; num++ {
		num_is_prime_factor := true
		for _, pf := range list_prime_factor {
			if num%pf == 0 {
				num_is_prime_factor = false
				break
			}
		}
		if num_is_prime_factor {
			return num
		}
	}
}
