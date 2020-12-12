package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type state int

const (
	floor state = iota
	empty
	occupied
)

func countOccupiedNeighbors(i int, j int, layout [][]state) int {
	countOccupied := 0
	// check all 8 possible around (i,j) directions from (i-1,j-1) to (i+1,j+1)
	for k := i - 1; k <= i+1; k++ {
		// ignore positions outside of the layout
		if k < 0 || k >= len(layout) {
			continue
		}
		for l := j - 1; l <= j+1; l++ {
			// ignore positions outside of the layout and (i,j) itself
			if l < 0 || l >= len(layout[0]) || (k == i && l == j) {
				continue
			}
			if layout[k][l] == occupied {
				countOccupied++
			}
		}
	}
	return countOccupied
}

func getNextState(i int, j int, layout [][]state, curState state) state {
	// floor stays constant
	if curState == floor {
		return floor
	}
	occupiedNeighbors := countOccupiedNeighbors(i, j, layout)
	// if an empty seat has no neighbors, it gets filled
	// if a seat is occupied and has 4 or more neighbors, it empties
	// otherwise it stays the same
	if curState == empty && occupiedNeighbors == 0 {
		return occupied
	} else if curState == occupied && occupiedNeighbors >= 4 {
		return empty
	}
	// remains unchanged
	return curState
}

func copyLayout(src [][]state) [][]state {
	dest := make([][]state, len(src))
	for i, states := range src {
		dest[i] = make([]state, len(states))
		for j, state := range states {
			dest[i][j] = state
		}
	}
	return dest
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run pt1.go <input_file>")
		os.Exit(1)
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	seatLayout := make([][]state, 0)
	scanner := bufio.NewScanner(file)
	lineNo := 0
	for scanner.Scan() {
		line := scanner.Text()
		curRow := make([]state, len(line))
		for i, char := range strings.Split(line, "") {
			switch char {
			case "L":
				curRow[i] = empty
			case "#":
				curRow[i] = occupied
			case ".":
				curRow[i] = floor
			default:
				fmt.Println("unexpected character on line", lineNo, "col", i)
				os.Exit(1)
			}
		}
		seatLayout = append(seatLayout, curRow)
		lineNo++
	}

	// create a workspace for building the next layout state
	nextSeatLayout := copyLayout(seatLayout)
	rounds := 0
	countOccupied := 0
	stable := false
	for !stable {
		// assume this iteration will be stable unless proven otherwise
		stable = true
		countOccupied = 0
		for i, line := range seatLayout {
			for j, curState := range line {
				nextState := getNextState(i, j, seatLayout, curState)
				if nextState != curState {
					stable = false
				}
				// count these here so we don't need another loop after stabilization
				if nextState == occupied {
					countOccupied++
				}
				nextSeatLayout[i][j] = nextState
			}
		}
		seatLayout = copyLayout(nextSeatLayout)
		rounds++
	}
	fmt.Printf("Took %d iterations to achieve stability with %d occupied seats.\n", rounds-1, countOccupied)
}
