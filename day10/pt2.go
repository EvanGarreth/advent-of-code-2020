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
		fmt.Println("Usage: go run pt2.go <input_file>")
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
	sort.Ints(joltages)

	// count the number of ways to reach the device with a valid adapter chain
	dp := make([]int, len(joltages))
	// starting joltage and first neighbor are only reachable one way
	dp[0] = 1
	dp[1] = 1
	// for each joltage, check the number of ways the previous 3 joltages were reachable and add that count
	//  to the current joltage IFF the previous joltage to be checked can bridge the current joltage
	// start at the first unknown value, index 2
	for i := 2; i < len(joltages); i++ {
		// reduce current number by 3 to check if it is useable by any of the 3 previous values
		numReach := joltages[i] - 3
		if numReach <= joltages[i-1] {
			dp[i] += dp[i-1]
		}
		if numReach <= joltages[i-2] {
			dp[i] += dp[i-2]
		}
		if i > 2 && numReach <= joltages[i-3] {
			dp[i] += dp[i-3]
		}
	}
	// the last dp value (index of the device adapter) is the number of valid ways to reach it
	fmt.Println("Number of valid permutations of adapters is:", dp[len(dp)-1])
}
