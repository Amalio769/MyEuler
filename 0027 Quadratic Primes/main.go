// Euler discovered the remarkable quadratic formula: n² + n + 41
// It turns out that the formula will produce 40 primes for the consecutive integer values 0<=n<=39.
// However, when n=40, 40²+40+41=40(40+1)+41 is divisible by 41, and certainly when n=41,
// 41²+41+41 is clearly divisible by 41.
// The incredible formula n²-79n+1601 was discovered, which produce 80 primes for the consecutive
// values 0<=n<=79. The product of the coefficients, -79 and 1601, is -126479.
// Considering quadratics of the form:
// n²+an+b, where |a| < 1000 and |b|<=1000
// where |n| is the modulus/absolute value of n
// e.g. |11| = 11 and |-4| = 4
// Find the product of the coefficients, a and b, for the quadratic expression that produces the
// maximum number of primes for consecutive values of n, starting with n=0.
// Anwser:
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// Para valor n=0 y el máximo valor de b (1000), el resultado debe ser un numero primo comprendido entre
	// 0 y 1000
	// Primero calculamos los numeros primos hasta 1000
	// Por lo tanto los posibles valores de b son el listado de numeros primos hasta 1000

	list_values_b := ListPrimesUntilNum(1001)
	fmt.Println(list_values_b)

	// Ahora para n=1, n²+an+b=1+a+b, por lo tanto para n=1 debe ser un primo hasta 2001
	list_result_n1 := ListPrimesUntilNum(2001)
	list_values_a := []int{}
	for b:= range list_values_b{
		for a:= -999;a<1000;a++{
			result:=1+a+b
			if result in list_result_n1
		}
	}

	fmt.Println("Computational time :", time.Since(start))
}

func ListPrimesUntilNum(num int) []int {
	var list_prime []int
	for i := 2; i < num; i++ {
		next_prime := NextPrime(list_prime)
		if next_prime < num {
			list_prime = append(list_prime, next_prime)
		} else {
			break
		}
	}
	return list_prime
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

func IsPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
