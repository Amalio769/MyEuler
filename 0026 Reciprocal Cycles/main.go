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
// Answer: 983
package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	max_recurring_cicle := 0
	max_number := 0
	for i := 2; i < 1000; i++ {
		result := LenRecurringCycleOneDividedNum(i)
		if result > max_recurring_cicle {
			max_recurring_cicle = result
			max_number = i
		}
	}
	fmt.Printf("For the number 1/%d, the result has a recurring cicle decimals of %d\n", max_number, max_recurring_cicle)

	fmt.Println("Computational time: ", time.Since(start))
}

func LenRecurringCycleOneDividedNum(num int) int {
	list_remainder := []int{}
	next_num := 1
	for {
		remainder := next_num % num
		if IsInSlice(list_remainder, remainder) || remainder == 0 {
			list_remainder = append(list_remainder, remainder)
			break
		} else {
			list_remainder = append(list_remainder, remainder)
			next_num = remainder * 10
		}
	}
	// fmt.Println(list_remainder)
	value := list_remainder[len(list_remainder)-1]
	first_ocurrence_value, err := FindIndex(list_remainder, value)
	if err != nil {
		return 0
	}
	last_ocurrence_value := len(list_remainder) - 1

	return last_ocurrence_value - first_ocurrence_value
}

func IsInSlice(slice []int, num int) bool {
	for _, value := range slice {
		if value == num {
			return true
		}
	}
	return false
}

func FindIndex(slice []int, num int) (int, error) {
	for idx, value := range slice {
		if value == num {
			return idx, nil
		}
	}
	return -1, errors.New("número no encontrado en el slice")
}
