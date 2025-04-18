// A Pythagorean triplet is a set of three natural numbers, a<b<c, for which,
// a²+b²=c²
// There exists exactly one Pythagorean triplet for which a+b+c=1000.
// Find the product abc.
// Answer:  31875000
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Computational time: ", time.Since(start))
}

func IsPythagoreanTriplet(a, b, c int) bool {
	if a*a+b*b == c*c {
		return true
	}
	return false
}
