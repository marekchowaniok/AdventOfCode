package days

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

type DBStructureNewFormat struct {
	ApplicationLogId string `csv:"ApplicationLogId"`
	UserName string `csv:"UserName"`
	Type string `csv:"Type"`
	Source string `csv:"Source"`
	IPAddress string `csv:"IPAddress"`
	Url string `csv:"Url"`
	Action string `csv:"Action"`
	Name string `csv:"Name"`
	Detail string `csv:"Detail"`
	StackTrace string `csv:"StackTrace"`
	ResponseTime string `csv:"ResponseTime"`
	DeletionState string `csv:"DeletionState"`
	CreatedOn string `csv:"CreatedOn"`
	ModifiedOn string `csv:"ModifiedOn"`
	CreatedBy string `csv:"CreatedBy"`
	ModifiedBy string `csv:"ModifiedBy"`
}

type YQoute struct {
	tokuchuNumbers []float64
	configurationNumbers []float64
	startTimes []time.Time
}

func YKDBParserNewFormat() {
	bytes, err := ioutil.ReadFile("/Users/marek/git/AdventOfCode/2020/inputs/YK-DEV-2020-09-01-short2.input")
	if err != nil {
		return
	}

	contents := string(bytes)
	split := strings.Split(contents, "\n")
	_ = split

	//var sum = 0.0
	//
	//var times []float64
	var startTime []time.Time

	const longForm = "1/2/2006 3:4:5 PM"
	recordsTokuchu_old := make([]map[string]int, len(split))
	recordsTokuchu := []float64{}
	recordsConfigurations := []float64{}
	//make(map[string]string),
	_ = recordsTokuchu_old

	for _,row := range split {
		//fmt.Println(row)
		if row == "" {
			break
		}
		splitColumns := strings.Split(row, "\t")
		ck := splitColumns[0]
		cv := strings.Split(splitColumns[0], ";")
		rv := strings.Split(cv[2], "$")
		time,_  := strconv.ParseFloat(rv[1],32)
		recordsTokuchu = append(recordsTokuchu,time )
		if strings.HasSuffix(ck,"YQuote.Tokuchu") {
			recordsTokuchu = append(recordsTokuchu,time )
		} else if strings.HasSuffix(ck,"YQuote.CreateConfiguration") {
			recordsConfigurations = append(recordsConfigurations, time)
		}


	}
	mmType := YQoute{

		tokuchuNumbers: recordsTokuchu,
		configurationNumbers: recordsConfigurations,
		startTimes: startTime,
	}

	fmt.Printf("UnSorted Tokuchu time:\t\t\n%v\n", mmType.tokuchuNumbers)

	fmt.Println(" ")
	fmt.Println(" ")

	fmt.Println("Start times:  ", mmType.startTimes)
	fmt.Println(" ")
	fmt.Println(" ")

	fmt.Println(" ", mmType.tokuchuNumbers)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	sort.Float64s(mmType.tokuchuNumbers) // sort

	fmt.Printf("Sorted numbers:\t\t\n%v\n", mmType.tokuchuNumbers)
	fmt.Println("------------------")
	fmt.Println("------------------")
	fmt.Printf("Minimum value:\t\t%v\n", mmType.GetMinValue())
	fmt.Println("------------------")
	//fmt.Printf("Range of values:\t%v\n", mmType.CalcRangeValues())
	//fmt.Println("------------------")
	fmt.Printf("Mean/AVG value:\t\t%v\n", mmType.CalcMean())
	fmt.Println("------------------")
	fmt.Printf("Median value:\t\t%v\n", mmType.CalcMedian())
	fmt.Println("------------------")
	fmt.Printf("Maximum value:\t\t%v\n", mmType.GetMaxValue())
	fmt.Println("------------------")
	fmt.Printf("Count of values:\t%v\n", mmType.count())
	fmt.Println("------------------")
	fmt.Println("------------------")
	fmt.Printf("API START Time difference:\t%v\n", mmType.GetMinMaxTime())
	//fmt.Println(" ")
	//fmt.Println(times)

	//fmt.Println(" ")
	//fmt.Println(sum)
}

func (mm YQoute) GetMinMaxTime() string {
	sort.Slice(mm.startTimes, func(i, j int) bool {
		return mm.startTimes[i].Before(mm.startTimes[j])
	})
	if mm.startTimes == nil {
		return ""
	}
	return mm.startTimes[0].String() + " -> " + mm.startTimes[len(mm.startTimes)-1].String()
}


func (mm *YQoute) GetMinValue() float64 {
	sort.Float64s(mm.tokuchuNumbers) // sort the numbers

	return mm.tokuchuNumbers[0]
}

// return the maximum value
func (mm *YQoute) GetMaxValue() float64 {
	sort.Float64s(mm.tokuchuNumbers) // sort the numbers

	return mm.tokuchuNumbers[len(mm.tokuchuNumbers)-1]
}

// calculate the range values
// last value - first value
func (mm *YQoute) CalcRangeValues() float64 {
	sort.Float64s(mm.tokuchuNumbers) // sort the numbers

	return mm.GetMaxValue() - mm.GetMinValue()
}

// calculate the "mean" value
// sum of all the values
// divided by its quantity
func (mm *YQoute) CalcMean() float64 {
	total := 0.0

	for _, v := range mm.tokuchuNumbers {
		total += v
	}

	// IMPORTANT: return was rounded!
	return math.Round(total / float64(len(mm.tokuchuNumbers)))
}



// calculate the "median" value
// if the total of numbers is odd
// 	takes the middle value
//
// if the total of numbers is even
//	calculate the "mean" of the middle two values
func (mm *YQoute) CalcMedian(n ...float64) float64 {
	sort.Float64s(mm.tokuchuNumbers) // sort the numbers

	mNumber := len(mm.tokuchuNumbers) / 2

	if mm.IsOdd() {
		return mm.tokuchuNumbers[mNumber]
	}

	return (mm.tokuchuNumbers[mNumber-1] + mm.tokuchuNumbers[mNumber]) / 2
}

// check if the total of numbers is
// odd or even
func (mm *YQoute) IsOdd() bool {
	if len(mm.tokuchuNumbers)%2 == 0 {
		return false
	}

	return true
}

func (mm *YQoute) count() int {
	return len (mm.tokuchuNumbers)
}