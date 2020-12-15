package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// the computer uses 36 bits, so use 64 bit ints to hold the info
type computer struct {
	setOnes  int64
	setZeros int64
	memory   map[int64]int64
}

func (c *computer) updateMask(newMask string) {
	// this mask will be AND'd with the value when writing to force 0s
	var setZeros int64 = 0xfffffffff
	// unset all 36 bits, will be OR'd with the value when writing to force 1s
	var setOnes int64 = 0
	for i, bit := range newMask {
		if bit == 'X' {
			continue
		}
		// calculate the number of bits to shift to the left since we iterate over the mask from high to low bits
		shift := len(newMask) - i - 1
		if bit == '0' {
			// unset the corresponding bit
			setZeros ^= 1 << shift
		} else if bit == '1' {
			// set the corresponding bit
			setOnes |= 1 << shift
		}
	}
	c.setOnes = setOnes
	c.setZeros = setZeros
}

func (c *computer) setBits(addr int64, val int64) {
	maskedVal := val
	// OR the value to force all 0's to 1's where masked
	maskedVal |= c.setOnes
	// AND the value to set all 1's to 0's where masked
	maskedVal &= c.setZeros
	c.memory[addr] = maskedVal
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

	// since addresses with unchanged values are considered 0, can use a map to hold only the ones that are changed
	memory := make(map[int64]int64, 0)
	c := &computer{0, 0, memory}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Index(line, "mask") != -1 {
			// the mask always starts at index 7
			index := 7
			c.updateMask(line[index:])
		} else {
			// the starting number for the address is always at index 4, so find the index of the terminating "]" for the end of the number
			right := strings.Index(line, "]")
			// convert to a 64 bit int since we use up to 36 bits
			addr, _ := strconv.ParseInt(line[4:right], 10, 64)
			// the starting point for the value is always + 2 from the equal sign and runs to the end of the string
			left := strings.Index(line, "=") + 2
			val, _ := strconv.ParseInt(line[left:], 10, 64)

			c.setBits(addr, val)
		}
	}

	var sum int64 = 0
	for _, val := range c.memory {
		sum += val
	}
	fmt.Println("Sum of all set memory values:", sum)
}
