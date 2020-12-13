package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	// earliest time we can board the bus
	departureTime, err := strconv.Atoi(lines[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	minDiff := math.MaxInt64
	minTime := 0
	for _, stime := range strings.Split(lines[1], ",") {
		// buses with an ID x are out of service
		if stime == "x" {
			continue
		}

		time, err := strconv.Atoi(stime)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// the amount of minutes we must wait for the first departure of this bus once we are ready
		// departureTime % time will give us the difference between the last departure before we are ready and when we are ready,
		// subtract that number from the bus' loop time to get the number of minutes we have to wait for the next departure once we are ready
		diff := abs(time - (departureTime % time))
		if diff < minDiff {
			minDiff = diff
			minTime = time
		}
	}

	fmt.Printf("Bus with closest departure to departure time of %d: %d.\n", departureTime, minTime)
	fmt.Printf("Bus ID * minutes waited: %d\n", minTime*minDiff)
}
