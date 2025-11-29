// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is
// 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.
// Answer:  906609
// Completed on Sun, 14 Aug 2022, 23:32

package main

import "fmt"

func main() {
	max_palindrome := 0
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			result := i * j
			if IsPalindrome(result) {
				if result > max_palindrome {
					max_palindrome = result
				}
			}
		}
	}
	fmt.Println(max_palindrome)
}

func IsPalindrome(n int) bool {
	return ReverseNumber(n) == n
}

func ReverseNumber(n int) int {
	reverse := 0
	temp := n
	for temp != 0 {
		reverse = (reverse * 10) + (temp % 10)
		temp = temp / 10
	}
	return reverse
}
