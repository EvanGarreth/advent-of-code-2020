package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operation int

const (
	nop operation = iota // do nothing
	acc                  // increase or decrease the program's accumulator by the argument value
	jmp                  // jump to another instruction relative to the current position by the argument value
)

// an instruction is an operation and an argument (a signed number)
type instruction struct {
	operation operation
	argument  int
}

// a program is a text file with one instruction per line, also contains a global accumulator
type program struct {
	instructions   []*instruction
	accumulator    int
	programCounter int
}

func (p *program) addInstruction(line string) error {
	splitString := strings.Split(line, " ")

	opName := splitString[0]
	var op operation
	switch opName {
	case "nop":
		op = nop
	case "acc":
		op = acc
	case "jmp":
		op = jmp
	default:
		return errors.New("unexpected operation provided")
	}

	argument := splitString[1]
	argVal, err := strconv.Atoi(argument)
	if err != nil {
		return err
	}

	instr := &instruction{op, argVal}
	p.instructions = append(p.instructions, instr)
	return nil
}

// load the instructions from the given file into memory
func (p *program) load(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		ok := p.addInstruction(line)
		if ok != nil {
			errMsg := fmt.Sprintf("Error on line %d: \"%s\"\n", lineNumber, err)
			return errors.New(errMsg)
		}
		lineNumber++
	}

	return nil
}

// make sure programCounter is in a valid state
func (p *program) canRun() bool {
	return p.programCounter >= 0 && p.programCounter < len(p.instructions)
}

// shouldn't be called outside of the program class since that checks for a valid programCounter
func (p *program) next() error {
	instruction := p.instructions[p.programCounter]
	switch instruction.operation {
	// add argument value to accumulator and increment program counter
	case acc:
		p.accumulator += instruction.argument
		p.programCounter++
	// add agument value to program counter to set up next instruction
	case jmp:
		nextProgramCounter := p.programCounter + instruction.argument
		if nextProgramCounter < 0 || nextProgramCounter > len(p.instructions) {
			errMSg := fmt.Sprintf("jmp instruction #%d would create an invalid program state", p.programCounter)
			return errors.New(errMSg)
		}
		p.programCounter = nextProgramCounter
	// do nothing (increment program counter to next instruction)
	case nop:
		p.programCounter++
	}

	return nil
}

// run the program, swapping a single jmp/nop at a time and backtracking if that didn't work.
// this will call itself and repeat until it runs
func (p *program) run(isCopy bool) (int, error) {
	if len(p.instructions) == 0 {
		return 0, errors.New("No program loaded to run")
	}

	// keep a map of the seen program counter states to track duplicate instructions
	seenInstructions := make(map[int]bool, len(p.instructions))
	for p.canRun() {
		pc := p.programCounter
		if seenInstructions[pc] {
			errMsg := fmt.Sprintf("Infinite loop detected @ pc %d", pc)
			return 0, errors.New(errMsg)
		}

		instr := p.instructions[pc]
		// if not already a copied program and is a swappable op
		// isCopy ensures only the first call to run() will make copies, meaning only 1 op will be swapped at a time
		if !isCopy && instr.operation != acc {
			// swap the op
			if instr.operation == jmp {
				instr.operation = nop
			} else {
				instr.operation = jmp
			}

			// save the current accumulator and PC to restore if this swap didn't work
			sAcc := p.accumulator
			sPC := p.programCounter
			res, err := p.run(true)
			// if the program didn't loop, return the result
			if err == nil {
				return res, nil
			}
			p.accumulator = sAcc
			p.programCounter = sPC

			// swap the op back for the next run
			if instr.operation == jmp {
				instr.operation = nop
			} else {
				instr.operation = jmp
			}
		}

		err := p.next()
		if err != nil {
			return 0, err
		}

		seenInstructions[pc] = true
	}

	return p.accumulator, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run pt2.go <input_file>")
		os.Exit(1)
	}

	fileName := os.Args[1]
	pg := program{make([]*instruction, 0), 0, 0}
	ok := pg.load(fileName)
	if ok != nil {
		fmt.Print(ok)
		os.Exit(1)
	}

	result, err := pg.run(false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("accumulator value at end of execution: %d\n", result)
}
