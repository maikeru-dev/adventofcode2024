package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func readInput(filename string) []byte {
	file, err := os.Open(filename) // Read access default

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	return data
}
func parseDataIntoRows(data []byte) [][]int {
	size := len(data)
	lineCounter := 0
	rows := make([][]int, 1000)

	// Dynamic row size, however each item
	for i := 0; i < size; i++ {
		valueCounter := 0
		row := make([]int, 10) // Say 10 max
		for {
			sizeCounter := 0
			currentValue := 0
			var charByte byte

			for {
				if i >= size {
					break
				}
				charByte = data[i]
				i++
				// Now attempt to read:
				if charByte == '\n' || charByte == '\r' || charByte == ' ' {
					break // Next row!
				}

				value, err := strconv.Atoi(string(charByte)) // this could be rewritten
				if err != nil {
					log.Fatal(err)
					return nil
				}

				if sizeCounter == 0 {
					currentValue = value
					sizeCounter++
					continue
				}

				currentValue = currentValue*10 + value
			}

			row[valueCounter] = currentValue
			valueCounter++

			if charByte == '\n' || charByte == '\r' {
				break
			}
			if i >= size {
				break
			}
		}
		rows[lineCounter] = row[:valueCounter]
		valueCounter = 0
		lineCounter++
	}

	return rows[:lineCounter]
}
func printRowCount(pairs []Pairs) {
	fmt.Printf("\n%d", len(pairs))
}
func main() {
	filename := "input.txt"
	file := readInput(filename)
	rows := convertRowsToPairs(parseDataIntoRows(file))
	rows = filterByDirection(rows)
	rows = filterByDifference(rows)

	for i := 0; i < len(rows); i++ {
		printRow(rows[i])
	}
	printRowCount(rows)
}
