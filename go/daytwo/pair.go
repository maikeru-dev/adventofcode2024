package main

import "fmt"

type Pair struct {
	a, b int
}
type Pairs struct {
	pairSet    []Pair
	faultCount int
}

func removeLink(index int, rows Pairs) Pairs {
	// Index refers to beginning index of the link
	// (index, next number)
	if index >= len(rows.pairSet) {
		return rows
	}

	rowsInt := make([]int, len(rows.pairSet)+1)
	for i := 0; i < len(rows.pairSet); i++ {
		if i >= index {
			rowsInt[i] = rows.pairSet[i+1].a
			if i == len(rows.pairSet) {
				// Last item, pull b instead.
				rowsInt[i+1] = rows.pairSet[i+1].b
				break
			}
			continue
		}
		rowsInt[i] = rows.pairSet[i].a
		if i == len(rows.pairSet) {
			// Last item, pull b instead.
			rowsInt[i+1] = rows.pairSet[i+1].b
			break
		}
	}

	return Pairs{newPairs, rows.faultCount}
}

// plural
func convertRowsToPairs(rows [][]int) []Pairs {
	rowsOfPairs := make([]Pairs, len(rows))

	for i := 0; i < len(rows); i++ {
		rowsOfPairs[i] = convertRowToPairs(rows[i])
	}

	return rowsOfPairs
}

// singular
func convertRowToPairs(row []int) Pairs {
	pairs := make([]Pair, len(row)-1)
	for i := 0; i < len(row)-1; i++ {
		pairs[i] = Pair{row[i], row[i+1]}
	}

	return Pairs{pairSet: pairs, faultCount: 0}
}
func checkDirection(pair Pair) int {
	if pair.a > pair.b {
		return 1
	}
	return -1
}

func filterByDirection(rows []Pairs) []Pairs {
	if len(rows) == 0 {
		return []Pairs{}
	}

	filteredRows := make([]Pairs, len(rows))
	rowCount := 0
	for i := 0; i < len(rows); i++ {
		result, row := checkDirectionOfPairs(rows[i])
		if result && row.faultCount == 1 {
			result, row = checkDirectionOfPairs(row)
		}
		if result {
			filteredRows[rowCount] = row
			rowCount++
		}
	}
	return filteredRows[:rowCount]
}

func checkDirectionOfPairs(pairs Pairs) (bool, Pairs) {
	if len(pairs.pairSet) == 0 {
		return true, pairs
	}
	// pool
	// Left -1, Right +1
	pairSet := pairs.pairSet
	directionPool := Pair{0, 0}
	for i := 0; i < len(pairSet); i++ {
		if checkDirection(pairSet[i]) == -1 {
			directionPool.a++
		} else {
			directionPool.b++
		}
	}

	var direction int
	if directionPool.a > directionPool.b { // majority a
		if directionPool.a == len(pairSet) { // fully left
			return true, pairs
		} else if directionPool.b > 1 { // right has some more than one
			return false, pairs
		}
		direction = -1
	} else if directionPool.a < directionPool.b { // majority b
		if directionPool.b == len(pairSet) { // fully right
			return true, pairs
		} else if directionPool.a > 1 { // left has some more than one
			return false, pairs
		}
		direction = 1
	} else {
		return false, pairs
	}
	pairs.faultCount++

	// for directionPool are off by exactly one, now identify bad index
	var index int
	for i := 0; i < len(pairSet); i++ {
		if directionPool.a == 1 { // In majority, increasing
			if checkDirection(pairSet[i]) != direction {
				index = i
				break
			}
		} else if directionPool.b == 1 { // In majority, decreasing
			if checkDirection(pairSet[i]) != direction {
				index = i + 1
				break
			}
		}
	}

	newRows := removeLink(index, pairs)

	return true, newRows
}

func difference(pair Pair) int {
	if pair.a > pair.b {
		return pair.a - pair.b
	}

	return pair.b - pair.a
}

func validatePairByDifference(pair Pair) bool {
	diff := difference(pair)
	return diff >= 1 && diff <= 3
}

// Adjusted for part 2
func validatePairsByDifference(pairs Pairs) bool {
	for i := 0; i < len(pairs.pairSet); i++ {
		if !validatePairByDifference(pairs.pairSet[i]) {
			if pairs.faultCount == 1 {
				return false
			}
			pairs.faultCount++
		}
	}

	return true
}

func filterByDifference(rows []Pairs) []Pairs {
	filteredRows := make([]Pairs, len(rows))
	rowCount := 0
	for i := 0; i < len(rows); i++ {
		if validatePairsByDifference(rows[i]) {
			filteredRows[rowCount] = rows[i]
			rowCount++
		}
	}

	return filteredRows[:rowCount]
}

func printRow(row Pairs) {
	output := "\n"
	if len(row.pairSet) < 1 {
		fmt.Print("\nEmpty output")
		return
	}
	for i := 0; i < len(row.pairSet); i++ {
		output += fmt.Sprintf("%d ", row.pairSet[i].a)
	}
	fmt.Print(output)
}
