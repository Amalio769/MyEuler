// Starting in the top left corner of a 2×2 grid, and only being able to move
// to the right and down, there are exactly 6 // routes to the bottom right
// corner.
// How many such routes are there through a 20×20 grid?
// Answer:  137846528820
// Grid Size:
// A 20×20 grid means you need to make exactly 40 moves (20 right and 20 down)
// to reach the bottom-right corner.
// Unique Paths:
// The problem reduces to arranging 2 "R" (right) moves and 20 "D" (down) moves
// in a sequence of 40 moves.
// The number of unique arrangements is given by the binomial coefficient:
// [ \binom{40}{20} = \frac{40!}{20! \cdot 20!}]
// The formula for binomial coeficient is:
// [ \binom{n}{k} = \frac{n!}{k! \cdot (n-k)!}]
package main

import (
	"fmt"
	"math/big"
)

func latticePaths(gridSize int) *big.Int {
	// Calculate (2n)!
	n := big.NewInt(int64(gridSize))
	k := new(big.Int).Mul(n, big.NewInt(2))
	kFactorial := new(big.Int).MulRange(1, k.Int64())

	// Calculate n!
	nFactorial := new(big.Int).MulRange(1, n.Int64())

	// Calculate (n!)^2
	nFactorialSquared := new(big.Int).Mul(nFactorial, nFactorial)

	// Calculate (2n)! / (n!)^2
	result := new(big.Int).Div(kFactorial, nFactorialSquared)

	return result
}

func main() {
	gridSize := 20
	fmt.Printf("Number of lattice paths through a %dx%d grid: %s\n", gridSize, gridSize, latticePaths(gridSize))
}
