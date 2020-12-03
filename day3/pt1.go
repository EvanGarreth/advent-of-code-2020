package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Slope given in the problem description
	const XDIR = 3
	const YDIR = 1
	// number of lines in input is also equal to max y dist
	// hardcoding this since the input will never change
	const NUMLINES = 323
	// width of each line, the map also repeats along the x-axis
	// also hardcoding since the input will never change
	const LINEWIDTH = 31
	const TREE = '#'

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	coordMap := make([]string, NUMLINES)
	lineNumber := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coordMap[lineNumber] = scanner.Text()
		lineNumber++
	}

	// Traverse the map, remembering that it repeats along the x-axis
	// Start at YDIR since line 0 is never visited
	xPos := 0
	yPos := YDIR
	treeCount := 0
	for yPos < NUMLINES {
		xPos = (xPos + XDIR) % LINEWIDTH
		if coordMap[yPos][xPos] == TREE {
			treeCount++
		}
		yPos += YDIR
	}

	fmt.Printf("Path of slope (%d, %d) will hit %d trees.\n", XDIR, YDIR, treeCount)
}
