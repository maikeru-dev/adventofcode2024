package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

var filename = "input.txt"

func readInput() []byte {
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

func slice(data []byte, index int, length int) []byte {
	return data[index:(index + length)]
}

func parseDataIntoLists(data []byte) ([]int, []int) {
	// 5 digit number, 3 character space, 5 digit number, 1 new line character
	var leftList, rightList []int

	size := len(data)
	lineLength := 5 + 3 + 5 + 1
	lineSize := size / lineLength

	leftList = make([]int, lineSize)
	rightList = make([]int, lineSize)

	for i := 0; i < lineSize; i++ { // per line

		leftNum, errL := strconv.Atoi(string(slice(data, (i*2)*5+(i*4), 5)))
		rightNum, errR := strconv.Atoi(string(slice(data, (i*2)*5+(5+3)+(i*4), 5)))

		if errL != nil {
			log.Fatal(errL)
		} else if errR != nil {
			log.Fatal(errR)
		}

		leftList[i] = leftNum
		rightList[i] = rightNum
	}

	return leftList, rightList
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("DayOne: ")

	data := readInput()

	leftList, rightList := parseDataIntoLists(data)
	sort.Ints(leftList)
	sort.Ints(rightList)

	total := 0
	for i := range leftList {
		distance := absDiffInt(leftList[i], rightList[i])

		total += distance
	}

	fmt.Printf("Total: %d", total)
}

