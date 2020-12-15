package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	numbers      map[int]int
	lastNum      int
	lastWasFirst bool
	round        int
	maxRounds    int
}

func (g *game) adjustNum() {
	num := g.lastNum
	age := 0
	// if we marked that the last number was seen for the first time, play a 0
	if g.lastWasFirst {
		age = 0
	} else {
		lastSeen := g.numbers[num]
		age = (g.round - 1) - lastSeen
	}

	// mark that the last num was played on the previous round
	// Don't mark on the round it is generated to preserve the previous round it was used on to calculate its age
	g.numbers[num] = g.round - 1
	// the age is the number we played this round so save it for the next iteration
	g.lastNum = age
	// if ok==true, this age has been played before so set the wasFirst flag accordingly
	_, ok := g.numbers[age]
	g.lastWasFirst = !ok

	g.round++
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run pt1.go <input_file> <rounds>")
		os.Exit(1)
	}

	roundLimit, _ := strconv.Atoi(os.Args[2])
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	game := &game{}
	game.numbers = make(map[int]int, 0)
	game.maxRounds = roundLimit
	game.lastNum = 0
	game.round = 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, val := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(val)
			game.numbers[num] = game.round
			game.round++
			game.lastNum = num
			game.lastWasFirst = true
		}
	}

	for game.round <= game.maxRounds {
		game.adjustNum()
	}

	fmt.Printf("%dth number is %d.\n", roundLimit, game.lastNum)
}
