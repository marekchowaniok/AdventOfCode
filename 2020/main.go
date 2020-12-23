package main

import (
	days "./days"
	"fmt"
	"os"
)

func main() {
	var day = "1"
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

	default:
		fmt.Println("Pick day between 1 and 24")
		days.One()
	}
}
