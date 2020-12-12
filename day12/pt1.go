package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type cardinalDir int

const (
	north cardinalDir = iota
	east
	south
	west
)

type turn int

const (
	left turn = iota
	right
)

type position struct {
	x int
	y int
}

type ship struct {
	curPos position
	facing cardinalDir
}

func (s *ship) move(dir cardinalDir, units int) {
	switch dir {
	case north:
		s.curPos.y += units
	case east:
		s.curPos.x += units
	case south:
		s.curPos.y -= units
	case west:
		s.curPos.x -= units
	}
}

func (s *ship) turn(turnDir turn, degrees int) {
	// degrees in the input file only take the values 90, 180, and 270
	// 180 means to turn in the opposite direction regardless of L or R,
	// a left turn at 90 degrees is equal to a right turn at 270,
	// and the same in reverse
	switch s.facing {
	case north:
		if degrees == 180 {
			s.facing = south
		} else if (turnDir == left && degrees == 90) || (turnDir == right && degrees == 270) {
			s.facing = west
		} else {
			s.facing = east
		}
	case east:
		if degrees == 180 {
			s.facing = west
		} else if (turnDir == left && degrees == 90) || (turnDir == right && degrees == 270) {
			s.facing = north
		} else {
			s.facing = south
		}
	case south:
		if degrees == 180 {
			s.facing = north
		} else if (turnDir == left && degrees == 90) || (turnDir == right && degrees == 270) {
			s.facing = east
		} else {
			s.facing = west
		}
	case west:
		if degrees == 180 {
			s.facing = east
		} else if (turnDir == left && degrees == 90) || (turnDir == right && degrees == 270) {
			s.facing = south
		} else {
			s.facing = north
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

	curPos := position{0, 0}
	ferry := ship{curPos, east}
	lineNo := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Each line is a single character for an action followed by an integer for units to move (or degrees if L or R)
		action := line[0]
		units, ok := strconv.Atoi(line[1:])
		if ok != nil {
			fmt.Printf("expected integer unit on line %d", lineNo)
			os.Exit(1)
		}

		switch action {
		case 'F':
			ferry.move(ferry.facing, units)
		case 'L':
			ferry.turn(left, units)
		case 'R':
			ferry.turn(right, units)
		case 'N':
			ferry.move(north, units)
		case 'E':
			ferry.move(east, units)
		case 'S':
			ferry.move(south, units)
		case 'W':
			ferry.move(west, units)
		default:
			fmt.Println("unexpected action on line", lineNo)
			os.Exit(1)
		}

		lineNo++
	}

	manhattanDist := abs(ferry.curPos.x) + abs(ferry.curPos.y)

	fmt.Printf("Manhattan distance is %d after sailing to (%d,%d)\n", manhattanDist, ferry.curPos.x, ferry.curPos.y)
}
