package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// number of rows of seats
const rowMax = 128

// number of columns of seats
const colMax = 8

type seatRange struct {
	lo int
	hi int
}

func bspSearch(input string) (int, int, error) {
	rowRange := seatRange{0, rowMax - 1}
	colRange := seatRange{0, colMax - 1}

	if len(input) != 10 {
		return -1, -1, errors.New("bspSearch expects input of length 10")
	}

	// first 7 chars are for rowRange, last determines whether to use lo or hi
	for _, char := range input[:6] {
		mid := rowRange.lo + (rowRange.hi-rowRange.lo)/2
		if char == 'F' {
			// take lower half of range
			rowRange.hi = mid
		} else if char == 'B' {
			// take upper half of range
			rowRange.lo = mid + 1
		} else {
			return -1, -1, errors.New("unexpected character in input string")
		}
	}
	// last 3 are for colRange, last determines whether to use lo or hi
	for _, char := range input[7 : len(input)-1] {
		mid := colRange.lo + (colRange.hi-colRange.lo)/2
		if char == 'L' {
			// take lower half of range
			colRange.hi = mid
		} else if char == 'R' {
			// take upper half of range
			colRange.lo = mid + 1
		} else {
			return -1, -1, errors.New("unexpected character in input string")
		}
	}

	// last row char tells us lo or hi
	row := -1
	if input[6] == 'F' {
		row = rowRange.lo
	} else if input[6] == 'B' {
		row = rowRange.hi
	} else {
		return -1, -1, errors.New("unexpected character in input string")
	}

	// last col char tells us lo or hi
	col := -1
	if input[9] == 'L' {
		col = colRange.lo
	} else if input[9] == 'R' {
		col = colRange.hi
	} else {
		return -1, -1, errors.New("unexpected character in input string")
	}

	return row, col, nil
}

func max(array []int) int {
	if len(array) == 0 {
		return 0
	}

	max := array[0]
	for _, val := range array[1:] {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	const ROWMULT = 8

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run pt1.go <input_file>")
		os.Exit(1)
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	seatIds := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			row, col, ok := bspSearch(line)
			if ok != nil {
				fmt.Println(ok)
				continue
			}
			// id formula from problem description
			id := row*ROWMULT + col
			seatIds = append(seatIds, id)
		}
	}
	fmt.Println("Highest seat id is ", max(seatIds))
}
