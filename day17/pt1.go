package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dimension struct {
	points map[string]bool
}

func (d *dimension) countActiveNeighbors(p string) int {
	activeNeighbors := 0
	xyz := strings.Split(p, ",")
	x, _ := strconv.Atoi(xyz[0])
	y, _ := strconv.Atoi(xyz[1])
	z, _ := strconv.Atoi(xyz[2])

	for i := z - 1; i <= z+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := x - 1; k <= x+1; k++ {
				if i == z && j == y && k == x {
					continue
				}
				if d.points[fmt.Sprintf("%d,%d,%d", k, j, i)] == true {
					activeNeighbors++
				}
			}
		}
	}
	return activeNeighbors
}

func (d *dimension) copy() *dimension {
	dest := &dimension{make(map[string]bool)}
	for point, state := range d.points {
		dest.points[point] = state
	}
	return dest
}

func main() {
	const CYCLES = 6
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

	pocketDimension := &dimension{make(map[string]bool, 0)}
	scanner := bufio.NewScanner(file)
	lineNo := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, state := range strings.Split(line, "") {
			point := fmt.Sprintf("%d,%d,%d", i, lineNo, 0)
			switch state {
			case ".":
				pocketDimension.points[point] = false
			case "#":
				pocketDimension.points[point] = true
			}
		}
		lineNo++
	}

	rounds := 1
	currentlyActive := 0
	nextDimensionState := pocketDimension.copy()
	for rounds <= CYCLES {
		currentlyActive = 0
		// keep the current structure within a window so parts don't get cut off
		for i := -rounds; i <= rounds; i++ {
			for j := -CYCLES - rounds; j <= CYCLES+rounds; j++ {
				for k := -CYCLES - rounds; k <= CYCLES+rounds; k++ {
					point := fmt.Sprintf("%d,%d,%d", k, j, i)
					state := pocketDimension.points[point]
					activeNeighbors := pocketDimension.countActiveNeighbors(point)
					nextState := false
					if state && (activeNeighbors == 2 || activeNeighbors == 3) {
						nextState = true
						currentlyActive++
					} else if !state && activeNeighbors == 3 {
						nextState = true
						currentlyActive++
					}

					nextDimensionState.points[point] = nextState
				}
			}
		}
		pocketDimension = nextDimensionState.copy()
		rounds++
	}

	fmt.Printf("After %d iterations, there are %d active cubes\n", rounds-1, currentlyActive)
}
