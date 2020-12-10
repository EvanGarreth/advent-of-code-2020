package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type passwordInfo struct {
	pos1      int
	pos2      int
	character string
	password  string
}

func parseLine(line string, lineNumber int) (*passwordInfo, error) {
	// format is "number-number letter: password"
	blocks := strings.Split(line, " ")
	positions := strings.Split(blocks[0], "-")
	lo, ok := strconv.Atoi(positions[0])
	if ok != nil {
		return nil, ok
	}
	hi, ok := strconv.Atoi(positions[1])
	if ok != nil {
		return nil, ok
	}

	result := &passwordInfo{}
	// the characters from the input are not zero indexed
	result.pos1 = lo - 1
	result.pos2 = hi - 1
	// blocks[1] will be a single char and a colon, so just take the single char
	result.character = string(blocks[1][0])
	result.password = blocks[2]

	return result, nil
}

func main() {
	file, err := os.Open("../input.txt")
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

		char1 := string(parsed.password[parsed.pos1])
		char2 := string(parsed.password[parsed.pos2])
		// Exactly one character has to be present in its spot for the password to be valid, so use XOR
		if (strings.Compare(char1, parsed.character) == 0) != (strings.Compare(char2, parsed.character) == 0) {
			fmt.Println("looking for", parsed.character, "found:", char1, char2, "in", parsed.password)
			validPaswords++
		}
	}

	fmt.Println("Number of valid passwords: ", validPaswords)
}
