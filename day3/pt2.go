package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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

	// Slopes given by problem description
	paths := [5][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	product := 1
	for _, path := range paths {
		xdir := path[0]
		ydir := path[1]
		// Traverse the map, remembering that it repeats along the x-axis
		// Start at YDIR since line 0 is never visited
		xPos := 0
		yPos := ydir
		treeCount := 0
		for yPos < NUMLINES {
			xPos = (xPos + xdir) % LINEWIDTH
			if coordMap[yPos][xPos] == TREE {
				treeCount++
			}
			yPos += ydir
		}
		fmt.Printf("Path of slope (%d, %d) will hit %d trees.\n", xdir, ydir, treeCount)
		product *= treeCount
	}

	fmt.Println("Product of trees hit in each path: ", product)
}
