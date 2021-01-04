package days

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Seven() {
	bytes, err := ioutil.ReadFile("/Users/marek/git/AdventOfCode/2020/inputs/day07.input")
	if err != nil {
		return
	}
	//split := sliceInput(bytes)
	contents := string(bytes)
	split := strings.Split(contents, "\n")

	graph := parseInput(split)
	fmt.Printf("%v\n", graph)

	bags := []string{}
	for bag := range graph {
		if traverseGraph(graph, bag) {
			bags = append(bags, bag)
		}

	}
	fmt.Printf("part1: %v", bags)
	fmt.Printf("part1: %d\n", len(bags))

	graph = parseInput(split)
	// subtract 1 for itself
	count := traversSubbags(graph, "shiny gold") - 1
	fmt.Printf("part2: %v\n", count)

}

func parseInput(input []string) map[string]map[string]int {
	graph := map[string]map[string]int{}
	for _, line := range input {
		split := strings.Split(line, " contain ")
		color := split[0][:strings.Index(split[0], " bags")]
		graph[color] = map[string]int{}

		for _, content := range strings.Split(split[1], ", ") {
			if content == "no other bags." {
				continue
			}
			parts := strings.Split(content, " ")
			num, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Printf("Error during parse %v", input)
			}
			graph[color][parts[1]+" "+parts[2]] = num
		}
	}
	return graph
}

func traverseGraph(graph map[string]map[string]int, entry string) bool {
	//for v := range  entry[]
	if _, ok := graph[entry]["shiny gold"]; ok {
		return true
	}
	for subBags := range graph[entry] {
		if traverseGraph(graph, subBags) {
			return true
		}
	}
	return false
}

func traversSubbags(graph map[string]map[string]int, entry string) int {
	bags := 1

	for subBag, count := range graph[entry] {
		res := traversSubbags(graph, subBag)
		fmt.Printf("bag: %v, count: %d, res: %d count*res: %d \n", subBag, count, res, count*res)
		bags += count * res
	}
	return bags
}
