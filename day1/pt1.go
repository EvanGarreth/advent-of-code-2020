package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	const TARGET = 2020

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var m = make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		m[value] = TARGET - value
	}

	// find the key that matches TARGET - value
	for _, val := range m {
		i, ok := m[val]
		if ok {
			fmt.Println("pair (", val, i, "), product: ", val * i)
			return
		}
	}
}