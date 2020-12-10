package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func validColor(val string) bool {
	allowedColors := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, color := range allowedColors {
		if val == color {
			return true
		}
	}
	return false
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
		// byr (Birth Year) - four digits; at least 1920 and at most 2002.
		case "byr":
			{
				byr, ok := strconv.Atoi(val)
				if ok != nil {
					return nil, errors.New("expected integer for field \"byr\"")
				}

				if byr < 1920 || byr > 2002 {
					return nil, errors.New("value for field \"byr\" out of range")
				}
				passp.byr = val
			}
		// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		case "iyr":
			{
				iyr, ok := strconv.Atoi(val)
				if ok != nil {
					return nil, errors.New("expected integer for field \"iyr\"")
				}

				if iyr < 2010 || iyr > 2020 {
					return nil, errors.New("value for field \"iyr\" out of range")
				}
				passp.iyr = val
			}
		// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		case "eyr":
			{
				eyr, ok := strconv.Atoi(val)
				if ok != nil {
					return nil, errors.New("expected integer for field \"eyr\"")
				}

				if eyr < 2020 || eyr > 2030 {
					return nil, errors.New("value for field \"eyr\" out of range")
				}
				passp.eyr = val
			}
		/*
			hgt (Height) - a number followed by either cm or in:
				If cm, the number must be at least 150 and at most 193.
				If in, the number must be at least 59 and at most 76.
		*/
		case "hgt":
			{
				cm := strings.Index(val, "cm")
				in := strings.Index(val, "in")
				if cm != -1 {
					realVal, ok := strconv.Atoi(val[0:cm])
					if ok != nil {
						return nil, ok
					}

					if realVal < 150 || realVal > 193 {
						return nil, errors.New("value for field \"hgt\" out of range")
					}
				} else if in != -1 {
					realVal, ok := strconv.Atoi(val[0:in])
					if ok != nil {
						return nil, ok
					}
					if realVal < 59 || realVal > 76 {
						return nil, errors.New("value for field \"hgt\" out of range")
					}

				} else {
					return nil, errors.New("expected value for field \"hgt\" to be in or cm")
				}
				passp.hgt = val
			}
		// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		case "hcl":
			{
				if val[0] != '#' {
					return nil, errors.New("expected field \"hcl\" to lead with \"#\"")
				}
				if len(val) != 7 {
					return nil, errors.New("expected length of field \"hcl\" to be 7 characters")
				}
				for _, char := range val[1:] {
					if !(unicode.IsDigit(char) || (char >= 'a' && char <= 'f')) {
						return nil, errors.New("expected only digits after \"#\" in field \"hcl\"")
					}
				}
				passp.hcl = val
			}
		//	ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		case "ecl":
			{
				if !validColor(val) {
					return nil, errors.New("unexpected eye color for field \"ecl\"")
				}
				passp.ecl = val
			}
		//	pid (Passport ID) - a nine-digit number, including leading zeroes.
		case "pid":
			{
				if len(val) != 9 {
					return nil, errors.New("expected length of field \"pid\" to be 9 characters")
				}
				for _, digit := range val {
					if !unicode.IsDigit(digit) {
						return nil, errors.New("expected only digits in field \"pid\"")
					}
				}
				passp.pid = val
			}
		// cid (Country ID) - ignored, missing or not.
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
		fmt.Println("Usage: go run pt2.go <input_file>")
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
			if ok != nil {
				fmt.Println(ok)
			} else {
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
		if ok != nil {
			fmt.Println(ok)
		} else {
			passports = append(passports, lastPassport)
		}
	}

	fmt.Println("Number of valid passports: ", len(passports))
}
