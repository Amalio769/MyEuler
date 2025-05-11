// A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:
// 1/2 = 0.5
// 1/3 = 0.(3)
// 1/4 = 0.25
// 1/5 = 0.2
// 1/6 = 0.1(6)
// 1/7 = 0.(142857)
// 1/8 = 0.125
// 1/9 = 0.(1)
// 1/10 = 0.1
//
// Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.
//
// Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
// Answer:
package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	start := time.Now()

	// re := regexp.MustCompile(`(\d+?)+`)
	// one := big.NewFloat(1)
	// for i := 2; i <= 100; i++ {
	// 	denominator := big.NewFloat(float64(i)).SetPrec(9000)
	// 	result := new(big.Float).Quo(one, denominator)
	// 	result_text := result.Text('f', 130)
	// 	match := re.FindStringSubmatch(result_text)
	// 	if len(match) > 1 {
	// 		fmt.Printf("Para 1/%d : %s", i, match)
	// 	}
	// }
	number := "0.0114942528735632183908045977011494252873563218390804597701149425287356321839080459770114942528735632183908045977"
	re := regexp.MustCompile(`([0-9]+?)+`)
	match := re.FindStringSubmatch(number)
	fmt.Println(match)

	fmt.Println("Computational time: ", time.Since(start))
}
