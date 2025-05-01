// 2pow(15) = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

// What is the sum of the digits of the number 2pow(1000)?
// Answer:  1366
package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	base := big.NewInt(2)
	exponent := big.NewInt(1000)
	result := new(big.Int).Exp(base, exponent, nil)
	strResult := result.String()
	sum_strResult := 0
	for _, n := range strResult {
		sum_strResult += int(n - '0')
	}
	fmt.Println(sum_strResult)
	fmt.Println("Computational time: ", time.Since(start))
}
