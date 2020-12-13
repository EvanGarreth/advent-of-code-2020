package main

import (
	"bufio"
	"fmt"
	"math"
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

type position struct {
	x int
	y int
}

type waypoint struct {
	pos    position
	facing cardinalDir
}

type ship struct {
	pos position
	wp  waypoint
}

func (wp *waypoint) move(dir cardinalDir, units int) {
	switch dir {
	case north:
		wp.pos.y += units
	case east:
		wp.pos.x += units
	case south:
		wp.pos.y -= units
	case west:
		wp.pos.x -= units
	}
}

func (s *ship) sail(times int) {
	diff := position{s.wp.pos.x - s.pos.x, s.wp.pos.y - s.pos.y}
	newFerryPos := position{s.pos.x + diff.x*times, s.pos.y + diff.y*times}
	newWpPos := position{newFerryPos.x + diff.x, newFerryPos.y + diff.y}

	s.pos = newFerryPos
	s.wp.pos = newWpPos
}

func (s *ship) rotateWaypoint(clockwise bool, degrees int) {
	// Translate waypoint so its relative to the point of origin
	// essentially, pretend the ship went back to (0,0) and adjust the waypoint accordingly
	rotationPoint := position{s.wp.pos.x - s.pos.x, s.wp.pos.y - s.pos.y}

	// convert degrees to radians since that's what the math lib uses
	angle := float64(degrees) * math.Pi / 180
	fsin, fcos := math.Sincos(angle)
	// can safely convert these to int now since the provided degrees will only give values of -1, 0 and 1
	sin := int(fsin)
	cos := int(fcos)

	// formula for rotating a point around the origin.
	// Save the x value for use in the calculation of y'
	temp := rotationPoint.x
	if clockwise {
		rotationPoint.x = rotationPoint.x*cos + rotationPoint.y*sin
		rotationPoint.y = rotationPoint.y*cos - temp*sin
	} else {
		rotationPoint.x = rotationPoint.x*cos - rotationPoint.y*sin
		rotationPoint.y = rotationPoint.y*cos + temp*sin
	}

	// translate the waypoint back so its relative to the ships coords
	s.wp.pos = position{s.pos.x + rotationPoint.x, s.pos.y + rotationPoint.y}
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

	// ferry always starts at (0,0)
	ferryPos := position{0, 0}
	// waypoiny starts 10 units east and 1 unit north of the ferry
	wpPos := position{10, 1}
	wp := waypoint{wpPos, east}
	ferry := ship{ferryPos, wp}
	lineNo := 0
	scanner := bufio.NewScanner(file)
	fmt.Println(ferry)
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
			ferry.sail(units)
		case 'L':
			ferry.rotateWaypoint(false, units)
		case 'R':
			ferry.rotateWaypoint(true, units)
		case 'N':
			ferry.wp.move(north, units)
		case 'E':
			ferry.wp.move(east, units)
		case 'S':
			ferry.wp.move(south, units)
		case 'W':
			ferry.wp.move(west, units)
		default:
			fmt.Println("unexpected action on line", lineNo)
			os.Exit(1)
		}
		lineNo++
	}

	manhattanDist := abs(ferry.pos.x) + abs(ferry.pos.y)

	fmt.Printf("Manhattan distance is %d after sailing to (%d,%d)\n", manhattanDist, ferry.pos.x, ferry.pos.y)
}
