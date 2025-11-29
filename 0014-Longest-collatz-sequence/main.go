// The following iterative sequence is defined for the set of positive integers:
// n → n/2 (n is even)
// n → 3n + 1 (n is odd)
// Using the rule above and starting with 13, we generate the following sequence:
// 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
// Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
// Which starting number, under one million, produces the longest chain?
// NOTE: Once the chain starts the terms are allowed to go above one million.
// Answer:  837799
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	number_with_longest_chain := 0
	max_chain := 0
	for i := 1; i < 1000000; i++ {
		temp_chain := LenghtSequenceCollatz(i)
		if temp_chain > max_chain {
			number_with_longest_chain = i
			max_chain = temp_chain
		}
	}
	fmt.Println(number_with_longest_chain)
	fmt.Println("Computational time: ", time.Since(start))
}

func LenghtSequenceCollatz(num int) int {
	if num <= 0 {
		return 0
	}
	var sequence []int
	next := 0
	sequence = append(sequence, num)
	for {
		last := sequence[len(sequence)-1]
		if last%2 == 0 {
			next = last / 2
		} else {
			next = 3*last + 1
		}
		sequence = append(sequence, next)
		if next == 1 {
			break
		}
	}
	return len(sequence)
}
