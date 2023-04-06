package days

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Eight() {
	bytes, err := ioutil.ReadFile("/Users/marek/git/AdventOfCode/2020/inputs/day08.input")
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
	count := traversSubbags(graph, "shiny gold") -1
	fmt.Printf("part2: %v\n", count)

}
//
//func parseInput2(input []string) map[int]map[string]int {
//	graph := map[int]map[string]int{}
//	for i, line := range input {
//		split := strings.Split(line, " ")
//		command := split[0]
//		graph[i] = map[string]int{}
//		graph[i][split[0]] = split[1]
//
//		for _, content := range strings.Split(split[1], ", ") {
//			if content == "no other bags." {
//				continue
//			}
//			parts := strings.Split(content, " ")
//			num, err := strconv.Atoi(parts[0])
//			if err != nil {
//				fmt.Printf("Error during parse %v", input)
//			}
//			graph[color][parts[1]+" "+parts[2]] = num
//		}
//	}
//	return graph
//}
