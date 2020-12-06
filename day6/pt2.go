package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run pt2.go <input_file>")
		os.Exit(1)
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	groupAnswers := make(map[string]int, 0)
	peopleInGroup := 0
	totalAnswers := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// blank line means a new group is starting, so record the current group's results and clear the old map
		if len(line) == 0 {
			// only count questions that everyone in a group provided an answer to
			allAnswered := 0
			for _, v := range groupAnswers {
				if v == peopleInGroup {
					allAnswered++
				}
			}
			totalAnswers += allAnswered
			groupAnswers = make(map[string]int, 0)
			peopleInGroup = 0
		} else {
			peopleInGroup++
			for _, char := range strings.SplitAfter(line, "") {
				groupAnswers[char]++
			}
		}
	}

	// grab the last group's count. Won't be counted above unless the input file ends with 2 newlines
	if len(groupAnswers) > 0 {
		allAnswered := 0
		for _, v := range groupAnswers {
			if v == peopleInGroup {
				allAnswered++
			}
		}
		totalAnswers += allAnswered
	}

	fmt.Println("Sum of answer counts", totalAnswers)
}
