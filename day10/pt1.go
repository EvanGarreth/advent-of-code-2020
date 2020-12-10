package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func max(array []int) int {
	if len(array) == 0 {
		return 0
	}
	max := array[0]
	for _, num := range array[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	const STARTINGJOLT = 0

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

	joltages := make([]int, 1)
	// add the starting joltage to the list
	joltages[0] = STARTINGJOLT
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		jolts, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		joltages = append(joltages, jolts)
	}
	// add the joltage of the device to the list, which is always + 3 of the max adapter joltage
	end := max(joltages) + 3
	joltages = append(joltages, end)

	// sort the list to get the correct adapter order
	// then look for differences of size 1 and 3 between direct neighbors
	sort.Ints(joltages)
	j1Diff := 0
	j3Diff := 0
	for i, num := range joltages[1:] {
		prev := joltages[i]
		if num-prev == 1 {
			j1Diff++
		} else if num-prev == 3 {
			j3Diff++
		}
	}
	fmt.Printf("After sorting, found %d differences of 1 and %d differences of 2. Product of these is: %d\n", j1Diff, j3Diff, j1Diff*j3Diff)
}
