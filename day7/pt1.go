package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string, error) {
	n := len(s)
	if n == 0 {
		return s, "", errors.New("Cannot pop from empty stack")
	}
	return s[:n-1], s[n-1], nil
}

type bagRule struct {
	name  string
	holds []string
}

func lineToBag(line string) (*bagRule, error) {
	tokens := strings.Split(line, " ")

	// shortest valid line is 7 tokens long
	if len(tokens) < 7 {
		return nil, errors.New("line is smaller than expected")
	}

	// name: <adj> <color>
	// line: <name> bags contain [(no other bags.) | (# <name> bags, )* (# <name> bags.)]
	name := strings.Join(tokens[:2], " ")
	holds := make([]string, 0)
	bag := &bagRule{name, holds}

	if strings.Index(line, "no other bags") == -1 {
		// each bag will be 4 tokens long, # adj color bag
		// if bags are present, will start at token index 4
		for i := 4; i < len(tokens); i += 4 {
			name = strings.Join(tokens[i+1:i+3], " ")
			bag.holds = append(bag.holds, name)
		}
	}

	return bag, nil
}

func (b bagRule) canHold(proposed string) bool {
	for _, member := range b.holds {
		if strings.Compare(member, proposed) == 0 {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run pt1.go <input_file> <bag>. A bag is an adjective and color separated by a space")
		os.Exit(1)
	}

	fileName := os.Args[1]
	startBag := strings.Join(os.Args[2:4], " ")
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	bagRules := make(map[string]*bagRule, 0)
	lineNumber := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bag, err := lineToBag(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		bagRules[bag.name] = bag
		lineNumber++
	}

	// get direct holders of the supplied bag
	s := make(stack, 0)
	seen := make(map[string]bool, 0)
	for _, bag := range bagRules {
		if bag.canHold(startBag) {
			s = stack.Push(s, bag.name)
			seen[bag.name] = true
		}
	}
	totalHolders := len(s)

	for len(s) != 0 {
		// declare current and err here so we don't shadow s with :=
		var current string
		var err error
		s, current, err = stack.Pop(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		seen[current] = true
		for _, bag := range bagRules {
			if bag.canHold(current) && !seen[bag.name] {
				s = stack.Push(s, bag.name)
				totalHolders++
			}
		}
	}

	fmt.Printf("Count of bags that can eventually hold \"%s\": %d\n", startBag, totalHolders)
}
