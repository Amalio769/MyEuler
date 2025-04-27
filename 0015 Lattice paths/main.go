// Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6
// routes to the bottom right corner.

// How many such routes are there through a 20×20 grid?
// Answer:  137846528820
package main

import (
	"fmt"
	"time"

	"github.com/Amalio769/mygolib/mymath"
)

func main() {
	start := time.Now()
	n := 10
	result := mymath.Factorial(n)
	fmt.Println(result)
	fmt.Println("Computational time: ", time.Since(start))
}
