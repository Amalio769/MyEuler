// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
// Answer:  6857
// Completed on Sun, 14 Aug 2022, 22:37

// algorithm FindPrimesEratosthenes(n):
//     // INPUT
//     //    n = an arbitrary number
//     // OUTPUT
//     //    prime numbers smaller than n
//
//     A <- an array of size n with boolean values set to true
//
//     for i <- 2 to sqrt(n):
//         if A[i] is true:
//             j <- i^2
//             while j <= n:
//                 A[j] <- false
//                 j <- j + i
//
//     return the indices of A corresponding to true

package main

import "fmt"

func main() {
	fmt.Println(PrimeFactor(600851475143))
}

func FindPrimesEratosthenes(n int64) []int64 {
	A := make([]bool, n+1)
	for k := range n + 1 {
		if k == 0 {
			A[k] = false
		} else {
			A[k] = true
		}
	}
}
