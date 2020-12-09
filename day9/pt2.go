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

func minMax(array []int) (int, int) {
	if len(array) == 0 {
		return 0, 0
	}

	max := array[0]
	min := array[0]
	for _, val := range array[1:] {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	return min, max
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
	targetNum := 0
	for _, num := range nums[:preamble] {
		window[num]++
	}

	// start with the first value outside the preamble
	for i, num := range nums[preamble:] {
		// check to see if a pair in the window sums to the new number, otherwise that's the number we need
		if !isValid(window, num) {
			fmt.Println("Found value with no sum pair in preamble", num)
			targetNum = num
			break
		}

		// i starts at 0, not preamble, so just decrement the value at i since that's what left the window
		left := nums[i]
		window[left]--
		// increment the value that just entered the window
		window[num]++
	}

	// now that we have the target num, do yet another loop looking for a sequence of numbers summing to it
	sum := 0
	// keep a map of sums mapping to index at which they belong, allows finding the sum in O(n) as opposed to O(n^2)
	sums := make(map[int]int, len(nums))
	for i := range nums {
		sum += nums[i]
		// a sequence i..j that sums to targetNum has been found if sum[j] - sum[i] == k,
		// so if the current sum is s[j], check if s[i] (== s[j] - k) exists in the map already
		if lo, ok := sums[sum-targetNum]; ok {
			min, max := minMax(nums[lo : i+1])
			fmt.Printf("Found sequence nums[%d:%d] summing to %d. Min/Max values in sequence: (%d, %d). Sum of min/max: %d\n", lo, i+1, targetNum, min, max, min+max)
			break
		}
		sums[sum] = i
	}
}
