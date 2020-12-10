package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isValid(window map[int]int, num int) bool {
	for key, val := range window {
		// if val is 0 the key isn't within the current window
		if val <= 0 {
			continue
		}
		// check to see if the element that would sum to num exists in the current window
		needed := num - key
		if window[needed] > 0 {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run pt1.go <input_file> <preamble length>")
		os.Exit(1)
	}

	fileName := os.Args[1]
	preamble, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	nums := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, num)
	}

	// use a map for O(1) lookup of values within the preamble
	// populate the map with the first values
	window := make(map[int]int, preamble)
	for _, num := range nums[:preamble] {
		window[num]++
	}

	for i, num := range nums[preamble:] {
		// check to see if the necessary values are in range
		if !isValid(window, num) {
			fmt.Println("Found value with no sum pair in preamble", num)
			os.Exit(0)
		}

		// i starts at 0, not preamble, so just decrement the value at i since that's what left the window
		left := nums[i]
		window[left]--
		// increment the value that just entered the window
		window[num]++
	}

	fmt.Println("All values have a matching sum pair in preamble")
}
