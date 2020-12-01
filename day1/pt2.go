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

	// make a list from the map keys to make iteration smoother
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	i = 0
	for i := range keys {
		j := i + 1
		for j = range keys {
			val := TARGET - keys[i] - keys[j]
			_, ok := m[val]
			if ok {
				fmt.Println("triplet (", keys[i], keys[j], val, "), product: ", keys[i] * keys[j] * val)
				return
			}
		}
	}
}