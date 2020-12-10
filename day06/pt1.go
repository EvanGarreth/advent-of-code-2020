package main

import (
	"bufio"
	"fmt"
	"os"
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
	}

	// go doesn't have a built-in set, so just use a map to booleans
	groupAnswers := make(map[string]bool, 0)
	totalAnswers := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// blank line means a new group is starting, so record the current group's results and clear the old map
		if len(line) == 0 {
			totalAnswers += len(groupAnswers)
			groupAnswers = make(map[string]bool, 0)
		} else {
			for _, char := range strings.SplitAfter(line, "") {
				groupAnswers[char] = true
			}
		}
	}

	// grab the last group's count. Won't be counted above unless the input file ends with 2 newlines
	if len(groupAnswers) > 0 {
		totalAnswers += len(groupAnswers)
	}

	fmt.Println("Sum of answer counts", totalAnswers)
}
