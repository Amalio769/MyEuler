// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

// What is the 10 001st prime number?
// Answer:  104743
// Completed on Mon, 15 Aug 2022, 15:54
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	number := 2
	var list_prime []int
	for {
		if IsPrimeFactor(number) {
			list_prime = append(list_prime, number)
			number++
			if len(list_prime) == 10001 {
				break
			}
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
