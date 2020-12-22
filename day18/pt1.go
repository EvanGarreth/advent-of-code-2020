package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type stack struct {
	s []interface{}
}

func (s *stack) push(item interface{}) {
	s.s = append(s.s, item)
}

func (s *stack) pop() interface{} {
	n := len(s.s)
	if n == 0 {
		return nil
	}

	res := s.s[n-1]
	s.s[n-1] = nil
	s.s = s.s[:n-1]
	return res
}

func (s *stack) peek() interface{} {
	n := len(s.s)
	if n == 0 {
		return nil
	}

	return s.s[n-1]
}

func (s *stack) isEmpty() bool {
	return len(s.s) == 0
}

func parseAndEval(line string, precedence bool) int {
	values := &stack{}
	operators := &stack{}
	// all tokens are single characters, so don't need to do any complicated scanning
	for _, char := range line {
		switch char {
		case ' ':
			continue
		case '(':
			operators.push(char)
		case ')':
			{
				// have a full parenthesized expression, so evaluate it
				for !operators.isEmpty() {
					top := operators.peek()
					if top == '(' {
						break
					}
					eval(operators, values)
				}
				_ = operators.pop()
			}
		case '+', '*':
			{
				// evaluate any ops we can on the stack before adding the new one
				for !operators.isEmpty() {
					top := operators.peek().(rune)
					// pt2 just adjusts precedence rules, so check that when adding a new op
					if top == ')' || top == '(' || (precedence && !hasGreaterPrecedence(top, char)) {
						break
					}
					eval(operators, values)
				}
				operators.push(char)
			}
		default:
			// assume we have a number since nothing else matched
			val, ok := strconv.Atoi(string(char))
			if ok != nil {
				fmt.Println(ok)
				os.Exit(1)
			}
			values.push(val)
		}
	}
	// eval anything that remains
	for !operators.isEmpty() {
		eval(operators, values)
	}
	return values.s[0].(int)
}

func hasGreaterPrecedence(curOp rune, nextOp rune) bool {
	// assume multiplication (lower precedence)
	curOpPrec := 0
	nextOpPrec := 0
	// adjust if either are addition
	if curOp == '+' {
		curOpPrec = 1
	}
	if nextOp == '+' {
		nextOpPrec = 1
	}
	return curOpPrec > nextOpPrec
}

func eval(operators *stack, values *stack) {
	op := operators.pop().(rune)
	r := values.pop().(int)
	l := values.pop().(int)
	res := 0
	switch op {
	case '*':
		res = l * r
	case '+':
		res = l + r
	default:
		fmt.Println("Unexpected op, \"", string(op), "\"")
		os.Exit(1)
	}
	values.push(res)
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

	sum1 := 0
	sum2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum1 += parseAndEval(line, false)
		sum2 += parseAndEval(line, true)
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}
