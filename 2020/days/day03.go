package days

import (
	"fmt"
	"io/ioutil"
)

func Three() {
	bytes, err := ioutil.ReadFile("/Users/marek/git/AdventOfCode/2020/inputs/day03.input")
	if err != nil {
		return
	}
	split := sliceInput(bytes)

	trees := make([][]bool, len(split))
	for i, s := range split {
		trees[i] = make([]bool, len(s))
		for j, c := range s {
			trees[i][j] = (c == '#')
		}
	}

	//hit := treeHits(trees, 3, 1)
	//hit := treeHits(trees, 1, 2)
	fmt.Printf("%d, %d, %d, %d, %d, %d, ", treeHits(trees, 1, 1), treeHits(trees, 3, 1), treeHits(trees, 5, 1), treeHits(trees, 7, 1), treeHits(trees, 1, 2),
		treeHits(trees, 1, 1)*treeHits(trees, 3, 1)*treeHits(trees, 5, 1)*treeHits(trees, 7, 1)*treeHits(trees, 1, 2))
	//fmt.Printf("Tree hits: %d \n", hit)
}

func treeHits(trees [][]bool, right int, down int) int {
	hit := 0
	for time := 0; time*down < len(trees); time++ {
		row := time * down
		column := (time * right) % len(trees[row])
		//fmt.Printf("row:%d column:%d => tree:%v\n", row, column, trees[row][column])
		if trees[row][column] {
			hit++
		}
	}
	return hit
}

/*
row, column
0,0
2,1
4,2
*/
