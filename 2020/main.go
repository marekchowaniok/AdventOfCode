package main

import (
	days "./days"
	"fmt"
	"os"
)

func main() {
	var day = "2b"
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
	case "2b":
		fmt.Println("Running day 02b")
		days.Twob()
	case "3":
		fmt.Println("Running day 03")
		days.Three()
	case "4":
		fmt.Println("Running day 04")
		days.Four()
	case "5":
		fmt.Println("Running day 05")
		days.Five()
	case "6":
		fmt.Println("Running day 06")
		days.Six()
	case "7":
		fmt.Println("Running day 07")
		days.Seven()
	case "8":
		fmt.Println("Running day 08")
		days.Eight()
	case "YKOI":
		fmt.Println("Running day YK")
		days.YKDBParser()
	case "YK":
		fmt.Println("Running day YK")
		days.YKDBParserNewFormat()

	default:
		fmt.Println("Pick day between 1 and 24")

	}
}
