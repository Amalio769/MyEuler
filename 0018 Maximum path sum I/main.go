// By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

//    3
//   7 4
//  2 4 6
// 8 5 9 3

// That is, 3 + 7 + 4 + 9 = 23.

// Find the maximum total from top to bottom of the triangle below:

//                             75
//                           95 64
//                         17 47 82
//                       18 35 87 10
//                     20 04 82 47 65
//                   19 01 23 75 03 34
//                 88 02 77 73 07 63 67
//               99 65 04 28 06 16 70 92
//             41 41 26 56 83 40 80 70 33
//           41 48 72 33 47 32 37 16 94 29
//         53 71 44 65 25 43 91 52 97 51 14
//       70 11 33 28 77 73 17 78 39 68 17 57
//     91 71 52 38 17 14 91 43 58 50 27 29 48
//   63 66 04 68 89 53 67 30 73 16 69 87 40 31
// 04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

// NOTE: As there are only 16384 routes, it is possible to solve this problem by trying every route. However, Problem 67, is the same challenge
// with a triangle containing one-hundred rows; it cannot be solved by brute force, and requires a clever method! ;o)
// Answer:  1074
// Algorithm:
// 1.-Start from the second-to-last row of the triangle.
// 2.-For each element in the row, add the maximum of the two adjacent numbers from the row below.
// 3.-Repeat this process for all rows until you reach the top of the triangle.
// 4.-The top element will contain the maximum total.
// This approach ensures that each number in the triangle is updated to represent the maximum sum possible from that point to the bottom.

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	triangle := [][]int{
		{75},
		{95, 64},
		{17, 47, 82},
		{18, 35, 87, 10},
		{20, 4, 82, 47, 65},
		{19, 1, 23, 75, 03, 34},
		{88, 2, 77, 73, 7, 63, 67},
		{99, 65, 4, 28, 6, 16, 70, 92},
		{41, 41, 26, 56, 83, 40, 80, 70, 33},
		{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
	}
	// check if triangle is correct
	if !checkTriangleIsOK(triangle) {
		fmt.Println("The triangle is not correct.")
		return
	}
	fmt.Println("Resultado:", maximumPathSumTriangle(triangle))
	fmt.Println("Computational time: ", time.Since(start))
}

func checkTriangleIsOK(triangle [][]int) bool {
	for i, row := range triangle {
		if len(row) != i+1 {
			return false
		}
	}
	return true
}

func maximumPathSumTriangle(triangle [][]int) int {
	result := triangle
	for {
		if lt := len(result); lt > 1 {
			result_sum := []int{}
			pos_row_sum := lt - 2
			len_row_sum := len(result[lt-2])
			for i := 0; i < len_row_sum; i++ {
				result_sum = append(result_sum, result[pos_row_sum][i]+Max2Values(result[pos_row_sum+1][i], result[pos_row_sum+1][i+1]))
			}
			result = append(result[0:lt-2], result_sum)
		} else {
			break
		}
	}
	return result[0][0]
}

func Max2Values(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
