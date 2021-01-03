package days

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Six() {
	bytes, err := ioutil.ReadFile("/Users/marek/git/AdventOfCode/2020/inputs/day06.input")
	if err != nil {
		return
	}
	//split := sliceInput(bytes)
	contents := string(bytes)
	split := strings.Split(contents, "\n\n")

	yesSum := 0
	for _, group := range split {
		//fmt.Printf("%v\n", group)

		group = strings.ReplaceAll(group, "\n", "")
		yes := make(map[rune]int)
		for _, r := range group {
			yes[r] = yes[r] + 1
		}
		yesSum += len(yes)
	}
	fmt.Printf("part1: %v\n", yesSum)

	yesSumPart2 := 0
	for _, group := range split {
		//group = strings.ReplaceAll(group, "\n", "")
		yes := make(map[rune]int)
		for _, r := range group {
			yes[r] = yes[r] + 1
		}
		people := yes[10] + 1
		if people == 1 {
			yesSumPart2 += len(yes)
		} else {
			for v := range yes {
				//fmt.Printf("%v", v)
				if yes[v] == people {
					yesSumPart2++
				}
			}
		}
	}
	fmt.Printf("part2: %v", yesSumPart2)
}
