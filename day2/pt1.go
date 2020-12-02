package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type passwordInfo struct {
	min       int
	max       int
	character string
	password  string
}

func parseLine(line string, lineNumber int) (*passwordInfo, error) {
	// format is "number-number letter: password"
	blocks := strings.Split(line, " ")
	minMax := strings.Split(blocks[0], "-")
	min, ok := strconv.Atoi(minMax[0])
	if ok != nil {
		return nil, ok
	}
	max, ok := strconv.Atoi(minMax[1])
	if ok != nil {
		return nil, ok
	}

	result := &passwordInfo{}
	result.min = min
	result.max = max
	// blocks[1] will be a single char and a colon, so just take the single char
	result.character = string(blocks[1][0])
	result.password = blocks[2]

	return result, nil
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	validPaswords := 0
	lineNumber := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsed, ok := parseLine(line, lineNumber)
		if ok != nil {
			fmt.Printf("Error parsing line %d: \"%s\", skipping.\n", lineNumber, line)
			continue
		}

		count := strings.Count(parsed.password, parsed.character)

		if count >= parsed.min && count <= parsed.max {
			validPaswords++
		}
	}

	fmt.Println("Number of valid passwords: ", validPaswords)
}
