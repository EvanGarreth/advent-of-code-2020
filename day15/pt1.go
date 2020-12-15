package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	numbers   []int
	prevNum   int
	round     int
	maxRounds int
}

func (g *game) play() {
	num := g.prevNum
	age := 0
	// 0 means that the previous round was the first time that number was played, so play a 0 this round
	if g.numbers[num] == 0 {
		age = 0
	} else {
		lastSeen := g.numbers[num]
		// prevNum was last seen in the previous round, and directly before that, the round stored in numbers[]
		age = (g.round - 1) - lastSeen
	}

	// mark that the last num was played on the previous round
	// Don't mark on the round it is generated to preserve the previous round it was used on to calculate its age
	g.numbers[num] = g.round - 1
	// the age is the number we played this round so save it for the next iteration
	g.prevNum = age
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
	// arbitrary large number that can hold the values when running input.txt to 30000000 rounds
	game.numbers = make([]int, 90000000)
	game.maxRounds = roundLimit
	game.prevNum = 0
	game.round = 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// add the starting numbers before the game is played.
		for _, val := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(val)
			game.numbers[num] = game.round
			game.prevNum = num
			game.round++
		}
	}

	for game.round <= game.maxRounds {
		game.play()
	}

	fmt.Printf("%dth number is %d.\n", roundLimit, game.prevNum)
}
