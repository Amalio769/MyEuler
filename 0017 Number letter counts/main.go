// If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19
// letters used in total.

// If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one
// hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
// Answer:  21124
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	digits := []string{
		"",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	teens := []string{
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}
	tens := []string{
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}
	list_number := []string{}
	for _, c := range digits {
		var century string
		if c == "" {
			century = ""
		} else {
			century = c + "hundred"
		}
		for _, n := range digits {
			if n == "" || c == "" {
				list_number = append(list_number, century+n)
			} else {
				list_number = append(list_number, century+"and"+n)
			}

		}
		for _, n := range teens {
			list_number = append(list_number, century+"and"+n)
		}
		for _, n := range tens {
			for _, k := range digits {
				list_number = append(list_number, century+"and"+n+k)
			}
		}
	}
	list_number = append(list_number, "onethousand")
	count_letter := 0
	lista_len := []int{}
	for _, i := range list_number {
		count_letter += len(i)
		lista_len = append(lista_len, len(i))
	}
	fmt.Println(list_number)
	fmt.Println(lista_len)
	fmt.Println(count_letter)
	fmt.Println("Computational time: ", time.Since(start))
}
