package days

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Three() {
	bytes, err := ioutil.ReadFile("/Users/marek/git/AdventOfCode/2020/inputs/day03.input")
	if err != nil {
		return
	}

	split := sliceInput(bytes)
	//contents := string(bytes)
	//split := strings.Split(contents, "\n")
	//split = split[:len(split)-1]

	matches1 := 0
	matches2 := 0
	for _, s := range split {
		parsed := strings.Split(s, " ")
		if len(parsed) != 3 {
			fmt.Println("Encountered bad line: ", split)
		}
		lenSpec := strings.Split(parsed[0], "-")
		if len(lenSpec) != 2 {
			fmt.Println("Encountered bad line: ", parsed)
		}
		char := rune(parsed[1][0])
		password := parsed[2]
		min, err := strconv.Atoi(lenSpec[0])
		max, err := strconv.Atoi(lenSpec[1])
		if err != nil {
			fmt.Printf("Failed to parse %s \n", parsed[0])
		}

		chars := make(map[rune]int)
		contains := false
		for i, c := range password {
			chars[c]++
			if (min == i+1 || max == i+1) && (c == char) {
				contains = !contains
			}
		}
		if contains {
			matches2++
		}

		if chars[char] >= min && chars[char] <= max {
			matches1++
		}
	}
	fmt.Printf("\n 1st matches1: %d \n", matches1)
	fmt.Printf("\n 2nd matches2: %d \n", matches2)
}
