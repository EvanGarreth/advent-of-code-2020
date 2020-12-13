package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	busIDs := make([]int, 0)
	for _, item := range strings.Split(lines[1], ",") {
		busID, err := strconv.Atoi(item)
		if err != nil {
			busIDs = append(busIDs, 0)
		} else {
			busIDs = append(busIDs, busID)
		}
	}

	timestamp := 1
	valid := false
	for !valid {
		step := 1
		// assume this iteration is valid
		valid = true

		// check the current timestamp candidate for a match on each bus and their offset
		for offset, busID := range busIDs {
			if busID == 0 {
				continue
			}

			if (timestamp+offset)%busID != 0 {
				valid = false
				break
			}

			// multiply the step by each busID that matches the pattern, as the current candidate timestamp + the product
			// of each matching busID is also the next timestamp that all of the matching buses will meet again
			step *= busID
		}

		if !valid {
			timestamp += step
		}
	}
	fmt.Printf("First valid timestamp: %d\n", timestamp)
}
