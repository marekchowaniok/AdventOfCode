package days

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var hexColor = regexp.MustCompile("#[0-9a-f]{6}$")
var pidDigits = regexp.MustCompile("^[0-9]{9}$")

func Four() {
	bytes, err := ioutil.ReadFile("/Users/marek/git/AdventOfCode/2020/inputs/day04.input")
	if err != nil {
		return
	}
	//split := sliceInput(bytes)
	contents := string(bytes)
	split := strings.Split(contents, "\n\n")

	//trees := make([][]bool, len(split))
	passports := make([]map[string]string, len(split))
	//make(map[string]string),
	_ = passports
	//valid := 0

	for i, s := range split {
		passports[i] = make(map[string]string)
		//trees[i] = make([]bool, len(s))
		s = strings.ReplaceAll(s, "\n", " ")
		keyVals := strings.Split(s, " ")
		for _, k := range keyVals {
			kv := strings.Split(k, ":")
			passports[i][kv[0]] = kv[1]
		}

	}
	valid, strict := isValid(passports)
	fmt.Printf("%d %d", valid, strict)
}

func isValid(passports []map[string]string) (int, int) {
	valid := 0
outer:
	for _, passport := range passports {
		for _, k := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
			if passport[k] == "" {
				continue outer
			}
		}
		valid++
	}

	strict := 0
	for _, passport := range passports {
		byr, err := strconv.Atoi(passport["byr"])
		if err != nil || byr < 1920 || byr > 2002 {
			continue
		}
		iyr, err := strconv.Atoi(passport["iyr"])
		if err != nil || iyr < 2010 || iyr > 2020 {
			continue
		}
		eyr, err := strconv.Atoi(passport["eyr"])
		if err != nil || eyr < 2020 || eyr > 2030 {
			continue
		}

		if len(passport["hgt"]) < 3 {
			continue
		}
		hgt, err := strconv.Atoi(passport["hgt"][:len(passport["hgt"])-2])
		units := passport["hgt"][len(passport["hgt"])-2:]
		if err != nil {
			continue
		}
		switch units {
		case "cm":
			if hgt < 150 || hgt > 193 {
				continue
			}
		case "in":
			if hgt < 59 || hgt > 76 {
				continue
			}
		default:
			continue
		}

		if len(passport["hcl"]) != 7 {
			continue
		}
		if !hexColor.MatchString(passport["hcl"]) {
			continue
		}

		switch passport["ecl"] {
		case "amb":
		case "blu":
		case "brn":
		case "gry":
		case "grn":
		case "hzl":
		case "oth":
		default:
			continue
		}

		if !pidDigits.MatchString(passport["pid"]) {
			continue
		}
		strict++
	}
	return valid, strict
}
