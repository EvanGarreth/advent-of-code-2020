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

func parseAndEval(line string) int {
	values := &stack{}
	operators := &stack{}
	for _, char := range line {
		switch char {
		case ' ':
			continue
		case '(':
			operators.push(char)
		case ')':
			{
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
				for !operators.isEmpty() {
					top := operators.peek()
					if top == ')' || top == '(' {
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
	for !operators.isEmpty() {
		eval(operators, values)
	}
	return values.s[0].(int)
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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum1 += parseAndEval(line)
	}
	fmt.Println(sum1)
}
