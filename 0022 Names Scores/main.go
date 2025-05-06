// Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names,
// begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value
// by its alphabetical position in the list to obtain a name score.

// For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the
// 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

// What is the total of all the name scores in the file?
// Answer:  871198282
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	// Abrir el fichero
	file, err := os.Open("/home/amalio/Documentos/GitHub/MyEuler/data/0022_names.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close() // Asegurarse de cerrar el archivo

	// Crear un lector CSV
	reader := csv.NewReader(file)
	reader.Comma = ',' // Especificar el delimitador (coma)

	// Leer todos los registros
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo TXT:", err)
		return
	}

	names := records[0]
	sort.Strings(names)

	result := 0
	for i, name := range names {
		result += (i + 1) * SumPosLetterInWord(name)
	}
	fmt.Println("The total of total word scores is: ", result)
	fmt.Println("Computational time: ", time.Since(start))
}

func SumPosLetterInWord(word string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	result := 0
	for _, letter := range word {
		result += strings.Index(alphabet, strings.ToLower(string(letter))) + 1
	}
	return result
}
