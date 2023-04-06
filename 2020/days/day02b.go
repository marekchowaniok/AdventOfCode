package days

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative path to the input file")

func Twob() {
	flag.Parse()
	bytes, err := ioutil.ReadFile("/Users/marek/git/AdventOfCode/2020/inputs/day02.input")
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]
	matches := 0
	for i, s := range split {
		parsed := strings.Split(s, " ")
		if len(parsed) != 3 {
			fmt.Printf("Encountered bad line %s\n",s)
		}

		lenSpec := strings.Split(parsed[0], "-")	// "3-12"
		min, err:= strconv.Atoi(lenSpec[0])
		max, err := strconv.Atoi(lenSpec[1])
		if err != nil {
			fmt.Printf("Encountered problem with min/max value %s\n",s)
		}
		char := parsed[1][:1]		//"n:"
		password := parsed[2] //"xxxxx"

		chars := make(map[rune]int)
		for _, c := range password {
			chars[c]++
		}
		if chars[char] >= min && chars[char] <= max {
			matches++
		}


		_, _, _, _, _ = max, min, char, password, i



	}
	fmt.Printf("matches: %s\n", matches)
}
