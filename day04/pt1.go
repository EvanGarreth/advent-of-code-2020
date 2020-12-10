package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type passport struct {
	byr string // Birth Year
	iyr string // Issue Year
	eyr string // Expiration Year
	hgt string // Height
	hcl string // Hair Color
	ecl string // Eye Color
	pid string // Passport ID
	cid string // Country ID
}

func buildPassport(buffer string) (*passport, error) {
	passp := &passport{}
	parts := strings.Split(buffer, " ")
	fieldsSet := 0
	for _, part := range parts {
		temp := strings.Split(part, ":")
		id := temp[0]
		val := temp[1]

		switch id {
		case "byr":
			passp.byr = val
		case "iyr":
			passp.iyr = val
		case "eyr":
			passp.eyr = val
		case "hgt":
			passp.hgt = val
		case "hcl":
			passp.hcl = val
		case "ecl":
			passp.ecl = val
		case "pid":
			passp.pid = val
		case "cid":
			passp.cid = val
		default:
			return nil, errors.New("unexpected id for struct field")
		}

		fieldsSet++
	}

	// Deny passports with 1 or more empty fields UNLESS the missing field is "cid"
	if fieldsSet < 8 && !(fieldsSet == 7 && len(passp.cid) == 0) {
		return nil, errors.New("invalid passport, missing required field")
	}
	return passp, nil
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
	}

	scanner := bufio.NewScanner(file)
	buffer := ""
	passports := make([]*passport, 0)
	for scanner.Scan() {
		line := scanner.Text()

		// each passport is separated by a blank line in the input file,
		// while key:value fields for a passport can be separated by spaces and newlines
		if len(line) == 0 {
			currentPassport, ok := buildPassport(buffer)
			if ok == nil {
				// silently ignore invalid passports
				passports = append(passports, currentPassport)
			}
			// clear buffer after every use
			buffer = ""
		} else {
			// If the buffer is empty, append the line. Otherwise add a space inbetween lines
			if len(buffer) == 0 {
				buffer = line
			} else {
				buffer += " " + line
			}
		}
	}

	// if the last line isn't a newline, there will still be a passport in the buffer
	if len(buffer) > 0 {
		lastPassport, ok := buildPassport(buffer)
		if ok == nil {
			passports = append(passports, lastPassport)
		}
	}

	fmt.Println("Number of valid passports: ", len(passports))
}
