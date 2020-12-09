package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
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

type innerBag struct {
	name  string
	count int
}

type bagRule struct {
	name  string
	holds []innerBag
}

func countChildrenHelper(bags map[string]*bagRule, bagName string, dp map[string]int) int {
	topBag := bags[bagName]

	// base case, if no children return 1 to denote a single bag
	if len(topBag.holds) == 0 {
		return 1
	}

	// if this bag has been counted before return the stored number
	if val, ok := dp[bagName]; ok {
		return val
	}

	// start the count at one to count the inner bag holding the others
	count := 1
	for _, bag := range topBag.holds {
		count += bag.count * countChildrenHelper(bags, bag.name, dp)
	}
	// store the end result in case its needed again
	dp[bagName] = count
	return count
}

func countChildren(bags map[string]*bagRule, bagName string) int {
	// use a map to hold the values of bags counted before
	dp := make(map[string]int, len(bags))
	// subtract by one to return the number of bags contained by the first bag, not the total bags present
	return countChildrenHelper(bags, bagName, dp) - 1
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
	holds := make([]innerBag, 0)
	bag := &bagRule{name, holds}

	if strings.Index(line, "no other bags") == -1 {
		// each bag will be 4 tokens long, # adj color bag
		// if bags are present, will start at token index 4
		for i := 4; i < len(tokens); i += 4 {
			name = strings.Join(tokens[i+1:i+3], " ")
			count, err := strconv.Atoi(tokens[i])
			if err != nil {
				return nil, errors.New("expected integer number of inner bags")
			}
			iBag := innerBag{name, count}
			bag.holds = append(bag.holds, iBag)
		}
	}

	return bag, nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run pt2.go <input_file> <bag>. A bag is an adjective and color separated by a space")
		os.Exit(1)
	}

	fileName := os.Args[1]
	startBag := strings.Join(os.Args[2:4], " ")
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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

	res := countChildren(bagRules, startBag)
	fmt.Printf("Count of bags that \"%s\" holds: %d\n", startBag, res)
}
