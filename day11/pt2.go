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

// not particularly elegant but it works. Might come back to this later
// the diagnonals can be made less confusing by using a step var and adjusting k and l within the loop instead
// or just make another function to determine i+x,j+x given a direction, with a const array of the 8 dirs?
func countOccupiedNeighbors(i int, j int, layout [][]state) int {
	countOccupied := 0
	// top
	for k := i - 1; k >= 0; k-- {
		if layout[k][j] != floor {
			if layout[k][j] == occupied {
				countOccupied++
			}
			break
		}
	}
	// bottom
	for k := i + 1; k < len(layout); k++ {
		if layout[k][j] != floor {
			if layout[k][j] == occupied {
				countOccupied++
			}
			break
		}
	}
	// left
	for l := j - 1; l >= 0; l-- {
		if layout[i][l] != floor {
			if layout[i][l] == occupied {
				countOccupied++
			}
			break
		}
	}
	// right
	for l := j + 1; l < len(layout[0]); l++ {
		if layout[i][l] != floor {
			if layout[i][l] == occupied {
				countOccupied++
			}
			break
		}
	}
	// top left
	for k, l := i-1, j-1; k >= 0 && l >= 0; {
		if layout[k][l] != floor {
			if layout[k][l] == occupied {
				countOccupied++
			}
			break
		}
		k--
		l--
	}
	// top right
	for k, l := i-1, j+1; k >= 0 && l < len(layout[0]); {
		if layout[k][l] != floor {
			if layout[k][l] == occupied {
				countOccupied++
			}
			break
		}
		k--
		l++
	}
	// bottom left
	for k, l := i+1, j-1; k < len(layout) && l >= 0; {
		if layout[k][l] != floor {
			if layout[k][l] == occupied {
				countOccupied++
			}
			break
		}
		k++
		l--
	}
	// bottom right
	for k, l := i+1, j+1; k < len(layout) && l < len(layout[0]); {
		if layout[k][l] != floor {
			if layout[k][l] == occupied {
				countOccupied++
			}
			break
		}
		k++
		l++
	}
	return countOccupied
}

func getNextState(i int, j int, layout [][]state, curState state) state {
	if curState == floor {
		return floor
	}
	occupiedNeighbors := countOccupiedNeighbors(i, j, layout)
	if curState == empty && occupiedNeighbors == 0 {
		return occupied
	} else if curState == occupied && occupiedNeighbors >= 5 {
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
	//fmt.Println(seatLayout)

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
