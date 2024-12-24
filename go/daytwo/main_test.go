package main

import (
	"testing"
)

func TestConsumeLink(t *testing.T) {
	// create pair list
	row := []int{1, 2, 3, 4, 5}
	secondRow := []int{2, 3, 4, 5}
	rowOfPairs := convertRowToPairs(row)
	secondRowOfPairs := convertRowToPairs(secondRow)

	newRowPair := removeLink(0, rowOfPairs)

	if !equalCompare(newRowPair, secondRowOfPairs) {
		t.Errorf("Failed to match!")
	}
}
func TestDirection(t *testing.T) {

	filename := "test_input.txt"
	file := readInput(filename)
	rows := convertRowsToPairs(parseDataIntoRows(file))
	rows = filterByDirection(rows)
	t.Errorf("%d rows", len(rows))
}

func equalCompare(rowA Pairs, rowB Pairs) bool {
	if len(rowA.pairSet) != len(rowB.pairSet) {
		return false
	}
	for i := 0; i < len(rowA.pairSet); i++ {

		if (rowA.pairSet[i].a == rowB.pairSet[i].a) && (rowA.pairSet[i].b == rowB.pairSet[i].b) {
			continue
		}
		return false
	}

	return true
}
