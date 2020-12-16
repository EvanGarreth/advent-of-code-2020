package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type section int

const (
	fieldRules section = iota
	myTicket
	nearbyTickets
)

type rule struct {
	lo int
	hi int
}

type field struct {
	name  string
	rules []*rule
}

func (f *field) fromString(line string) {
	re := regexp.MustCompile(`([a-zA-Z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)`)
	parts := re.FindStringSubmatch(line)
	f.name = parts[1]
	rule1 := &rule{}
	rule1.lo, _ = strconv.Atoi(parts[2])
	rule1.hi, _ = strconv.Atoi(parts[3])
	rule2 := &rule{}
	rule2.lo, _ = strconv.Atoi(parts[4])
	rule2.hi, _ = strconv.Atoi(parts[5])

	f.rules = append(f.rules, rule1)
	f.rules = append(f.rules, rule2)
}

func (f *field) isValid(val int) bool {
	for _, rule := range f.rules {
		if val >= rule.lo && val <= rule.hi {
			return true
		}
	}
	return false
}

type ticket struct {
	values []int
	fields []*field
}

func (t *ticket) fromCSV(csv string) {
	for _, num := range strings.Split(csv, ",") {
		val, _ := strconv.Atoi(num)
		t.values = append(t.values, val)
	}
}

func (t *ticket) invalidValues() []int {
	invalidValues := make([]int, 0)
	for _, val := range t.values {
		matchesOne := false
		for _, field := range t.fields {
			if field.isValid(val) {
				matchesOne = true
				break
			}
		}
		if !matchesOne {
			invalidValues = append(invalidValues, val)
		}
	}
	return invalidValues
}

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

	curSection := fieldRules
	fields := make([]*field, 0)
	tickets := make([]*ticket, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// check for these strings so we can skip to the line containing the values
		if line == "your ticket:" {
			curSection = myTicket
			continue
		} else if line == "nearby tickets:" {
			curSection = nearbyTickets
			continue
		}

		if len(line) <= 0 {
			continue
		}

		switch curSection {
		case fieldRules:
			{
				field := &field{}
				field.fromString(line)
				fields = append(fields, field)
			}
		case myTicket, nearbyTickets:
			{
				ticket := &ticket{}
				ticket.fromCSV(line)
				ticket.fields = fields
				tickets = append(tickets, ticket)
			}
		}
	}

	sum := 0
	legalTickets := make([]*ticket, 0)
	// first ticket is my ticket and should not be included
	for _, ticket := range tickets[1:] {
		invalidValues := ticket.invalidValues()
		if len(invalidValues) == 0 {
			legalTickets = append(legalTickets, ticket)
		} else {
			// pt1 needs the sum of all non-legal values
			for _, val := range ticket.invalidValues() {
				sum += val
			}
		}
	}
	fmt.Printf("Sum of invalid fields: %d\n", sum)

	// I'm not very proud of the below code, but it works for the given input
	// TODO make this less convoluted?

	positions := make([]*field, len(fields))
	candidates := make(map[*field]map[int]bool, len(fields))
	// for each field
	for _, field := range fields {
		// iterate over each value column
		for i := range tickets[0].values {
			matchesAll := true
			// looking for an index that all values match this field
			for _, ticket := range legalTickets {
				if !field.isValid(ticket.values[i]) {
					matchesAll = false
					break
				}
			}
			// if so, mark this value column as a candidate for this field
			if matchesAll {
				if len(candidates[field]) == 0 {
					candidates[field] = make(map[int]bool)
				}
				candidates[field][i] = true
			}
		}
	}

	// the below loop assumes that for each iteration, (at least) one candidate will have a single
	// possible column that it covers and will never have to make a choice and branch
	stable := false
	for !stable {
		stable = true

		for k, v := range candidates {
			// find the candidate with only one possible column
			if len(v) == 1 {
				// is there a better way to get the only k:v in a map?
				for k2 := range v {
					// set the field to its column in positions
					positions[k2] = k
					stable = false
					// delete the position from all other candidates
					for c2 := range candidates {
						delete(candidates[c2], k2)
					}
				}
			}
		}
	}

	// pt2 want's the product of all values with a field starting with "departure"
	product := 1
	for i, field := range positions {
		if len(field.name) >= 9 && field.name[:9] == "departure" {
			product *= tickets[0].values[i]
		}
	}
	fmt.Println("Product of values of fields starting with \"departure\":", product)
}
