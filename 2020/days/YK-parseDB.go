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

type DBStructure struct {
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

type MeanMedian struct {
	numbers []float64
	startTimes []time.Time
}

func YKDBParser() {
	bytes, err := ioutil.ReadFile("/Users/marek/git/AdventOfCode/2020/inputs/YK-DEV-2020-09-01-short.input")
	if err != nil {
		return
	}
	//var datas []DBStructure
	//gocsv.UnmarshalBytes(bytes, &datas)
	//
	////print(datas)
	//for _,data := range datas {
	//	if data.Name == "CheckOrderInstruction Api" {
	//		fmt.Println("",data.Detail)
	//		fmt.Println("", data)
	//		fmt.Println(" ", )
	//	}
	//}
	//split := sliceInput(bytes)
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	_ = split

	var sum = 0.0

	var times []float64
	var startTime []time.Time

	const longForm = "1/2/2006 3:4:5 PM"

	for _,row := range split {
		//fmt.Println(row)
		parsed := strings.Split(row, "---")
		apiTime := strings.Split(parsed[1], ":")
		timee,_  := strconv.ParseFloat(apiTime[2],32)
		times = append(times, timee)
		//s = append(s, 9, 10)
		sum += timee
		//fmt.Println("Start times: ")
		//fmt.Print(parsed[5],", ")
		t, _ := time.Parse(longForm, parsed[5])
		if t.Year() == 1 {
			t, _ = time.Parse(longForm, parsed[6])
		}
		startTime = append(startTime, t)
		//fmt.Println(timee)
		//fmt.Println(parsed)
	}
	mmType := MeanMedian{
		numbers: times,
		startTimes: startTime,
	}

	fmt.Printf("Sorted numbers:\t\t\n%v\n", mmType.numbers)

	fmt.Println(" ")
	fmt.Println(" ")

	fmt.Println("Start times:  ", mmType.startTimes)
	fmt.Println(" ")
	fmt.Println(" ")

	fmt.Println(" ", mmType.numbers)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	sort.Float64s(mmType.numbers) // sort

	fmt.Printf("Sorted numbers:\t\t\n%v\n", mmType.numbers)
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

func (mm MeanMedian) GetMinMaxTime() string {
	sort.Slice(mm.startTimes, func(i, j int) bool {
		return mm.startTimes[i].Before(mm.startTimes[j])
	})

	return mm.startTimes[0].String() + " -> " + mm.startTimes[len(mm.startTimes)-1].String()
}


func (mm *MeanMedian) GetMinValue() float64 {
	sort.Float64s(mm.numbers) // sort the numbers

	return mm.numbers[0]
}

// return the maximum value
func (mm *MeanMedian) GetMaxValue() float64 {
	sort.Float64s(mm.numbers) // sort the numbers

	return mm.numbers[len(mm.numbers)-1]
}

// calculate the range values
// last value - first value
func (mm *MeanMedian) CalcRangeValues() float64 {
	sort.Float64s(mm.numbers) // sort the numbers

	return mm.GetMaxValue() - mm.GetMinValue()
}

// calculate the "mean" value
// sum of all the values
// divided by its quantity
func (mm *MeanMedian) CalcMean() float64 {
	total := 0.0

	for _, v := range mm.numbers {
		total += v
	}

	// IMPORTANT: return was rounded!
	return math.Round(total / float64(len(mm.numbers)))
}



// calculate the "median" value
// if the total of numbers is odd
// 	takes the middle value
//
// if the total of numbers is even
//	calculate the "mean" of the middle two values
func (mm *MeanMedian) CalcMedian(n ...float64) float64 {
	sort.Float64s(mm.numbers) // sort the numbers

	mNumber := len(mm.numbers) / 2

	if mm.IsOdd() {
		return mm.numbers[mNumber]
	}

	return (mm.numbers[mNumber-1] + mm.numbers[mNumber]) / 2
}

// check if the total of numbers is
// odd or even
func (mm *MeanMedian) IsOdd() bool {
	if len(mm.numbers)%2 == 0 {
		return false
	}

	return true
}

func (mm *MeanMedian) count() int {
	return len (mm.numbers)
}