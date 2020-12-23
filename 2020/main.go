package main

import (
	days "./days"
	"fmt"
	"os"
)

func main() {
	var day = "2"
	if len(os.Args) > 1 {
		day = os.Args[1]
	}
	whichDay(day)
}

func whichDay(day string) {
	switch day {
	case "1":
		fmt.Println("Running day 01")
		days.One()
	case "2":
		fmt.Println("Running day 02")
		days.Two()
	default:
		fmt.Println("Pick day between 1 and 24")

	}
}
