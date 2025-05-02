// You are given the following information, but you may prefer to do some research for yourself.
// 1 Jan 1900 was a Monday.
// Thirty days has September,
// April, June and November.
// All the rest have thirty-one,
// Saving February alone,
// Which has twenty-eight, rain or shine.
// And on leap years, twenty-nine.
// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
// Answer: 171
// Completed on Fri, 2 May 2025, 23:03
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// Contador de domingos
	sundaysCount := 0

	// Iterar sobre los años del siglo XX (1901 a 2000)
	for year := 1901; year <= 2000; year++ {
		// Iterar sobre los meses del año
		for month := time.January; month <= time.December; month++ {
			// Crear una fecha para el primer día del mes
			date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
			// Verificar si el día de la semana es domingo
			if date.Weekday() == time.Sunday {
				sundaysCount++
			}
		}
	}

	fmt.Printf("Número de domingos que cayeron en el primer día del mes: %d\n", sundaysCount)
	fmt.Println("Computational time: ", time.Since(start))
}
