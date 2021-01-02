package days

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func Five() {
	bytes, err := ioutil.ReadFile("/Users/marek/git/AdventOfCode/2020/inputs/day05.input")
	if err != nil {
		return
	}
	//split := sliceInput(bytes)
	contents := string(bytes)
	split := strings.Split(contents, "\n")

	seatID := 0.0
	seatIDMap := []int{}
	for _, line := range split {
		if len(line) == 0 {
			continue
		}
		//fmt.Printf("%v\n", line)
		row, _ := searchPlace(0, 127, line[:len(line)-3])
		column, _ := searchPlace(0, 7, line[len(line)-3:])
		currentSeat := (row * 8) + column
		seatIDMap = append(seatIDMap, currentSeat)
		seatID = math.Max(seatID, float64(currentSeat))
		//fmt.Printf("row: %d column: %d seatID: %f\n", row, column, seatID)
	}
	fmt.Printf("seatID: %f\n", seatID)
	sort.Ints(seatIDMap)
	prev := 0
	for i, v := range seatIDMap {
		if v != prev+1 && v != 0 && i > 3 {
			fmt.Printf("problematic: %v\n", v-1)
		}
		prev = v
	}
	//fmt.Printf("seatidMap: %v", seatIDMap)
}

func searchPlace(min, max int, ticket string) (int, int) {
	mid := (max - min) / 2
	mid = min + mid

	letter := ticket[:len(ticket)-(len(ticket)-1)]
	//fmt.Printf("%d %d %d\t %v\n", min, mid, max, letter)
	if len(ticket) == 1 {
		if ticket == "F" || ticket == "L" {
			return min, min
		} else {
			return max, max
		}
	}
	switch letter {
	case "F":
		return searchPlace(min, mid, ticket[len(ticket)-(len(ticket)-1):])
	case "L":
		return searchPlace(min, mid, ticket[len(ticket)-(len(ticket)-1):])
	case "B":
		return searchPlace(mid+1, max, ticket[len(ticket)-(len(ticket)-1):])
	case "R":
		return searchPlace(mid+1, max, ticket[len(ticket)-(len(ticket)-1):])
	}
	return mid, mid
}
